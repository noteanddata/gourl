package main

import "net/http"

func getSingle(url string, headerFilePath string) (*http.Response, error) {
	request, err := createGetRequestWithHeader(url, headerFilePath)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	return client.Do(request)
}
