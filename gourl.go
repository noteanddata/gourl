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
  post bool
  filePath string
}

type BatchCallResult struct {
  success int
  failure int
}

func run() {
  url := os.Args[len(os.Args)-1]
  var opt Options
  flag.BoolVar(&opt.printHeader, "h", false, "print header")
  flag.BoolVar(&opt.post, "p", false, "http post")
  flag.IntVar(&opt.number, "n", 1, "number of request")
  flag.IntVar(&opt.concurrent, "c", 1, "concurrent request")
  flag.StringVar(&opt.filePath, "f", "", "post file path")
  flag.Parse()

  //fmt.Printf("%+v \n", opt)
  if opt.post {
    dopost(url, opt)
  } else {
    doget(url, opt)  
  }
  
}

func dopost(url string, opt Options) {
  content_type := "application/x-www-form-urlencoded"
  postSingle(url, content_type, opt.filePath, opt.printHeader)
}



func doget(url string, opt Options) error {
  if opt.number == 1 {
      return getonePrint(url, opt.printHeader)
  } else {
    
    
      remain := opt.number % opt.concurrent
      
      t1 := time.Now()
        
      successCount := make(chan int)
      failureCount := make(chan int)
      timeMilliSeconds := make(chan int64)
      for i := 0; i < opt.concurrent; i++ {
        count := opt.number / opt.concurrent
        if i < remain {
          count++
        }
        go getn(url, count, successCount, failureCount, timeMilliSeconds)
      }
  
      totalSuccess, totalFailure, sumTimeMilliSeconds := 0, 0, int64(0)
      for i := 0; i < opt.concurrent; i++ {
        totalSuccess += <- successCount
        totalFailure += <- failureCount
        sumTimeMilliSeconds += <- timeMilliSeconds
      }
      
      totalTimeMilliSeconds := time.Now().Sub(t1).Nanoseconds() / 1000.0 / 1000.0
      avgTimeMilliSeconds := sumTimeMilliSeconds / int64(opt.number)
      
      qps := int64(opt.number) * 1000.0 / totalTimeMilliSeconds 
      successRatio := totalSuccess * 100.0 / opt.number
      fmt.Println("concurrent=", opt.concurrent, ",totalSuccess=", totalSuccess, ", totalFailure=", totalFailure, ", success ratio=", successRatio, "%")
      fmt.Println("total time(ms)=", totalTimeMilliSeconds, ", qps=", qps, ", avgTime(ms)=", avgTimeMilliSeconds)
      
      return nil
  }   
} 
 
func getn(url string, count int, successCount chan int, failureCount chan int, timeMilliSeconds chan int64)  {  
  success, failure := 0,0
  start := time.Now()
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
  timeMilliSeconds <- time.Now().Sub(start).Nanoseconds() / 1000.0 / 1000.0
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