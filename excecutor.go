package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// RequestInfo contains the information needed to build the request
type RequestInfo struct {
	url            string
	contentType    string
	headerFilePath string
	postFilePath   string
}

// RequestExecutor will execute a request
type RequestExecutor interface {
	execute(requestInfo RequestInfo) (*http.Response, error)
}

// GetRequestExecutor will execute a GET request
type GetRequestExecutor struct{}

// PostRequestExecutor will execute a POST request
type PostRequestExecutor struct{}

func (getRequestExecutor GetRequestExecutor) execute(requestInfo RequestInfo) (*http.Response, error) {
	return getSingle(requestInfo.url, requestInfo.headerFilePath)
}

func (postRequestExecutor PostRequestExecutor) execute(requestInfo RequestInfo) (*http.Response, error) {
	return postSingle(requestInfo.url, requestInfo.contentType, requestInfo.postFilePath, false, requestInfo.headerFilePath)
}

func executeN(count int, requestInfo RequestInfo, requestExecutor RequestExecutor, callStatistic chan<- CallStatistic) {
	success, failure := 0, 0
	start := time.Now()
	for i := 0; i < count; i++ {
		resp, err := requestExecutor.execute(requestInfo)
		if err != nil {
			failure++
			continue
		}

		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				failure++
				continue
			}

			if len(bodyBytes) >= 0 {
				success++
			}
		}
	}

	callStatistic <- CallStatistic{success, failure, time.Now().Sub(start).Nanoseconds() / 1000.0 / 1000.0}
}

func createRequestExecutor(opt Options) RequestExecutor {
	if opt.post {
		return PostRequestExecutor{}
	}
	return GetRequestExecutor{}
}

func createRequestInfo(url string, opt Options) RequestInfo {
	return RequestInfo{
		url:            url,
		contentType:    opt.contentType,
		headerFilePath: opt.headerFilePath,
		postFilePath:   opt.postFilePath,
	}
}

func concurrentExecute(url string, opt Options) {
	remain := opt.number % opt.concurrent
	requestExecutor := createRequestExecutor(opt)
	requestInfo := createRequestInfo(url, opt)
	callStatisticChan := make(chan CallStatistic)

	t1 := time.Now()

	for i := 0; i < opt.concurrent; i++ {
		count := opt.number / opt.concurrent
		if i < remain {
			count++
		}
		go executeN(count, requestInfo, requestExecutor, callStatisticChan)
	}

	totalSuccess, totalFailure, sumTimeMilliSeconds := 0, 0, int64(0)

	callStatistic := <-callStatisticChan

	for i := 0; i < opt.concurrent; i++ {
		totalSuccess += callStatistic.successCount
		totalFailure += callStatistic.failureCount
		sumTimeMilliSeconds += callStatistic.timeMilliSeconds
	}

	totalTimeMilliSeconds := time.Now().Sub(t1).Nanoseconds() / 1000.0 / 1000.0
	avgTimeMilliSeconds := sumTimeMilliSeconds / int64(opt.number)

	qps := int64(opt.number) * 1000.0 / totalTimeMilliSeconds
	successRatio := totalSuccess * 100.0 / opt.number
	fmt.Println("concurrent=", opt.concurrent, ",totalSuccess=", totalSuccess, ", totalFailure=", totalFailure, ", success ratio=", successRatio, "%")
	fmt.Println("total time(ms)=", totalTimeMilliSeconds, ", qps=", qps, ", avgTime(ms)=", avgTimeMilliSeconds)
}
