package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func printResponseDetail(printHeader bool, resp *http.Response) error {
	fmt.Println(resp.Status)

	if printHeader {
		for header, value := range resp.Header {
			fmt.Println(header, value)
		}
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return err
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}
	return nil

}
