package main

import "net/http"
import "os"
import "fmt"
import "io/ioutil"

func postSingle(url string, content_type string, filePath string, printHeader bool) error {
    file, err := os.Open(filePath)
    if err != nil {
      fmt.Println(err)
      return err
    }
    resp, err := http.Post(url, content_type, file)
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