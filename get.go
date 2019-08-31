package main

import (
	"fmt"
	"net/http"
)

func getSingle(url string, headerFilePath string) (*http.Response, error) {
	request, err := createGetRequestWithHeader(url, headerFilePath)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	return client.Do(request)
}

// get one url and print result
func getOnePrint(url string, headerFilePath string, printHeader bool) error {
	resp, err := getSingle(url, headerFilePath)

	if err != nil {
		fmt.Println(err)
		return err
	}

	err = printResponse(printHeader, resp)
	return err
}
