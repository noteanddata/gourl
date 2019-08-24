package main

import "net/http"
import "os"
import "fmt"
import "io/ioutil"

const (
	headerContentType = "Content-Type"
)

func postSingle(url string, contentType string, postFilePath string, printHeader bool) error {
	fmt.Println("postFilePath=", postFilePath)
	file, err := os.Open(postFilePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	request, err := http.NewRequest("POST", url, file)
	if err != nil {
		fmt.Println(err)
		return err
	}

	request.Header.Add(headerContentType, contentType)
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
