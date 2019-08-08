package main
import "fmt"
import "os"
import "bufio"
//import "strings"

const key_body_type ="body.type"

type PostFormConfig struct {
    postData map[string][]string
}

type PostConfig struct {
  isForm bool 
  postFormConfig PostFormConfig
}

type PostConfigParser interface {
  ParseConfig(postConfigFile string) (PostConfig, error)
}

type PostConfigParserImpl struct {
  
}

func (postConfigParser PostConfigParserImpl) ParseConfig(postConfigFile string) (PostConfig, error) {
  var ret PostConfig 
  file, err := os.Open(postConfigFile)
  if err != nil {
    fmt.Println("failed to open postConfigFile=", postConfigFile)
    return ret, nil
  }
  
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    fmt.Println(line)
    
  }
  if err := scanner.Err(); err != nil {
    fmt.Println(err)
  }
  return ret, nil
}  