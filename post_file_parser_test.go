package main

import "testing"
import "fmt"

func TestParseConfig_(t *testing.T) {
  file := "./resources/post_form_sample1.txt"
  postConfigParserImpl := PostConfigParserImpl{}
  postConfig, err := postConfigParserImpl.ParseConfig(file)
  fmt.Println(err)
  fmt.Println(postConfig)
}