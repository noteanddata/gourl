package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	headerContentType = "Content-Type"
)

func postSingle(url string, contentType string,
	postFilePath string, printHeader bool,
	headerFilePath string) error {

	file, err := os.Open(postFilePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	request, err := http.NewRequest("POST", url, file)
	if err != nil {
		fmt.Println(err)
		return err
	}

	request.Header.Add(headerContentType, contentType)

	if headerFilePath != "" {
		file, err := os.Open(headerFilePath)
		if err != nil {
			fmt.Println(err)
			return err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			arr := strings.Split(line, ":")
			request.Header.Add(arr[0], arr[1])
		}
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return err
	}

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
