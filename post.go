package main

import (
	"fmt"
	"net/http"
)

const (
	headerContentType = "Content-Type"
)

func postSingle(url string, contentType string,
	postFilePath string, printHeader bool,
	headerFilePath string) error {
	request, file, err := createPostRequest(url, postFilePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if file != nil {
		defer file.Close()
	}

	request, err = addHeaders(request, contentType, headerFilePath)
	if err != nil {
		fmt.Println(err)
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = printResponse(printHeader, resp)
	return err
}
