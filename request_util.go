package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func createPostRequest(url string, postFilePath string) (*http.Request, *os.File, error) {
	file, err := os.Open(postFilePath)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	// we can not close in this method because the actual read is outside
	// defer file.Close()

	request, err := http.NewRequest("POST", url, file)
	if err != nil {
		fmt.Println(err)
		return nil, file, err
	}
	return request, file, nil
}

func createGetRequestWithHeader(url string, headerFilePath string) (*http.Request, error) {
	request, err := createGetRequest(url)
	if err != nil {
		return nil, err
	}
	addHeaders(request, "", headerFilePath)

	return request, nil
}

func createGetRequest(url string) (*http.Request, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return request, nil
}

func addHeaders(request *http.Request, contentType string, headerFilePath string) (*http.Request, error) {
	if len(contentType) > 0 {
		request.Header.Add(headerContentType, contentType)
	}

	// default to gourl user agent
	request.Header.Add("User-Agent", "gourl")

	if headerFilePath == "" {
		return request, nil
	}

	file, err := os.Open(headerFilePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, ":")
		request.Header.Add(arr[0], arr[1])
	}

	return request, nil
}
