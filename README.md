# gourl
a simple url utility written by http://www.noteanddata.com for practicing go

# sample usage 

## simple get 
```
gourl  http://www.example.com
```

## with header
```
gourl -h http://www.example.com

```

## support n number of http get 
```
gourl -n 5 http://www.example.com

```

## support n number with c concurrent calls of http get

```
gourl -n 1000 -c 100 http://www.example.com
totalSuccess= 987 , totalFailure= 13 , time(ms)= 2398

```

## support http post 
```
gourl -p -pf ./resources/post_form_sample1.txt -h http://localhost:8080/generate_short_url

```

## support http post with content type header 
```
gourl -p -pf ./resources/post_sample.json -ct application/json -h http://localhost:8080/

```