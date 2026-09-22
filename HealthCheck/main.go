package main

import (
	"fmt"
	"net/http"
	"time"
)

type Result struct {
	URL    string
	Status string
	Error  error
}

func checkServer(url string, ch chan Result) {
	res, err := http.Get(url)
	if err != nil {
		ch <- Result{
			URL:    url,
			Status: "is down!!!",
			Error:  err,
		}
		return
	}
	ch <- Result{
		URL:    url,
		Status: "is Running...",
		Error:  err,
	}
	defer res.Body.Close()
}

func main() {
	ch := make(chan Result)
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.wrongurl.com",
	}

	for _, url := range urls {
		go checkServer(url,ch)
	}
	start:=time.Now()

	for range urls{
		res:=<-ch
		if res.Error!=nil{
			fmt.Println(res.URL,res.Status,"Error: ",res.Error)
			continue
		}
		fmt.Println("Server status: ",res.Status)
	}
	fmt.Println("All server check completed: ",time.Since(start))
}
