package main
import "os"
import "fmt"
import "io/ioutil"
import "net/http"
import "flag"

func main() {
  run()
}

type Options struct {
  printHeader bool
}

func run() {
  url := os.Args[len(os.Args)-1]
  var opt Options
  flag.BoolVar(&opt.printHeader, "h", false, "print header")
  flag.Parse()

  //fmt.Printf("%+v \n", opt)

  get(url, opt)
}

func get(url string, opt Options) error {
  resp, err := http.Get(url)
  if err != nil {
    fmt.Println(err)
    return err
  }
  // print status code 
  fmt.Println(resp.Status)
  
  if opt.printHeader {
    
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