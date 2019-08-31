package main

import "os"

import "flag"

const (
	headerContentType         = "Content-Type"
	contentTypeFormUrlencoded = "application/x-www-form-urlencoded"
)

func main() {
	run()
}

// Options contains gourl options
type Options struct {
	printHeader    bool
	number         int
	concurrent     int
	post           bool
	contentType    string
	postFilePath   string
	headerFilePath string
}

// CallStatistic contains information about statistics
type CallStatistic struct {
	successCount     int
	failureCount     int
	timeMilliSeconds int64
}

func run() {
	url := os.Args[len(os.Args)-1]
	var opt Options
	flag.BoolVar(&opt.printHeader, "h", false, "print header")
	flag.BoolVar(&opt.post, "p", false, "http post")
	flag.IntVar(&opt.number, "n", 1, "number of request")
	flag.IntVar(&opt.concurrent, "c", 1, "concurrent request")
	flag.StringVar(&opt.postFilePath, "pf", "", "post file path")
	flag.StringVar(&opt.contentType, "ct", "", "Content-Type header")
	flag.StringVar(&opt.headerFilePath, "hf", "", "header file path")
	flag.Parse()

	//fmt.Printf("%+v \n", opt)
	if opt.post {
		doPost(url, opt)
	} else {
		doGet(url, opt)
	}
}

func doPost(url string, opt Options) {
	if opt.number == 1 {
		contentType := contentTypeFormUrlencoded // default value
		if opt.contentType != "" {
			contentType = opt.contentType
		}
		postSingle(url, contentType, opt.postFilePath, opt.printHeader, true, opt.headerFilePath)
	} else {
		concurrentExecute(url, opt)
	}
}

func doGet(url string, opt Options) error {
	if opt.number == 1 {
		return getOnePrint(url, opt.headerFilePath, opt.printHeader)
	}
	concurrentExecute(url, opt)
	return nil
}
