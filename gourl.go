package main
import "os"
import "fmt"
import "io/ioutil"
import "net/http"
import "flag"
import "time"


func main() {
  run()
}

type Options struct {
  printHeader bool
  number int
  concurrent int
}

type BatchCallResult struct {
  success int
  failure int
}

func run() {
  url := os.Args[len(os.Args)-1]
  var opt Options
  flag.BoolVar(&opt.printHeader, "h", false, "print header")
  flag.IntVar(&opt.number, "n", 1, "number of request")
  flag.IntVar(&opt.concurrent, "c", 1, "concurrent request")
  flag.Parse()

  //fmt.Printf("%+v \n", opt)

  doget(url, opt)
}

func doget(url string, opt Options) error {
  if opt.number == 1 {
      return getonePrint(url, opt.printHeader)
  } else {
    
    
      remain := opt.number % opt.concurrent
      
      t1 := time.Now()
        
      successCount := make(chan int)
      failureCount := make(chan int)
      for i := 0; i < opt.concurrent; i++ {
        count := opt.number / opt.concurrent
        if i < remain {
          count++
        }
        go getn(url, count, successCount, failureCount)
      }
  
      totalSuccess, totalFailure := 0, 0
      for i := 0; i < opt.concurrent; i++ {
        totalSuccess += <- successCount
        totalFailure += <- failureCount
      }
      
      timeMilliSeconds := time.Now().Sub(t1).Nanoseconds() / 1000.0 / 1000.0
      
      fmt.Println("totalSuccess=", totalSuccess, ", totalFailure=", totalFailure, ", time(ms)=", timeMilliSeconds)
      
      return nil
  }   
} 
 
func getn(url string, count int, successCount chan int, failureCount chan int)  {  
  success, failure := 0,0
  
  for i := 0; i < count; i++ {
    resp, err := http.Get(url)
    if err != nil {
      failure++
      continue
    }
    
    if resp.StatusCode == http.StatusOK {
      bodyBytes, err := ioutil.ReadAll(resp.Body)
      if err != nil {
        failure++
        continue
      }
      
      if len(bodyBytes) >= 0 {
        success++
      }  
    }
  }
  
  successCount <- success
  failureCount <- failure
}

// get one url and print result 
func getonePrint(url string, printHeader bool) error {
  resp, err := http.Get(url)
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