package main

import (
	"fmt"
	"net/http"
)

func postSingle(url string, contentType string,
	postFilePath string, printHeader bool, printResponse bool,
	headerFilePath string) (*http.Response, error) {
	request, file, err := createPostRequest(url, postFilePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if file != nil {
		defer file.Close()
	}

	request, err = addHeaders(request, contentType, headerFilePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if printResponse {
		err = printResponseDetail(printHeader, resp)
	}

	return resp, err
}
