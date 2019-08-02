package main
import "os"
import "fmt"
import "io/ioutil"
import "net/http"

func main() {
  run()
}

func run() {
  url := os.Args[1]
  get(url)
}

func get(url string) error {
  resp, err := http.Get(url)
  if err != nil {
    fmt.Println(err)
    return err
  }
  fmt.Println(resp.Status)
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