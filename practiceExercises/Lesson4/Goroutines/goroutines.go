//This is a practice in goroutines (go's concurrency primative)
//Concurrency is the composition of independently executed processes
//get content type of sites
package main

import (
	"fmt"
	"net/http"
	"sync"
)

//function, returnType, that accepts a url string
func returnType(url string) {

	//response, error are the results of an http.Get request
	resp, err := http.Get(url)

	//if the error is not null
	if err != nil {

		//print and return the error
		fmt.Printf("error: %s\n", err)
		return
	}

	//close the body of the response at the end of the function
	defer resp.Body.Close()

	//variable, ctype, reads header info
	ctype := resp.Header.Get("content-type")

	//prints out header info
	fmt.Printf("%s -> %s\n", url, ctype)
}

func main() {

	//variable, url, contains a slice of urls
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			returnType(url)
			wg.Done()
		}(url)
	}
	wg.Wait()
}
