package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("network error")

func main() {
	c := make(chan error)
	var results = make(map[string]string)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}
	for _, url := range urls {
		result := "OK"
		err := <-c
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string, c chan<- error) {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil {
		c <- err
	} else if resp.StatusCode >= 400 {
		fmt.Println(resp.StatusCode)
		c <- errRequestFailed
	} else {
		c <- nil
	}

}
