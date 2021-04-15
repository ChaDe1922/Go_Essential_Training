// write a function that gets a URL and returns the value of Content-Type response HTTP header
// The function should return an error if it can't perform a GET request
//Use net/http Get function to make an HTTP call
// resp, err := http.Get("https:/golang.org")
// use resp.Header.Get to get the value of a header
//resp.Header.Get("Content-Length")
//make sure the response body is closed
//resp.Body.Close()

package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	resp, err := http.Get(url)

	//in go, we always check for the error!
	if err != nil {
		return "", err
	}

	defer resp.Body.Close() //make sure the response body is closed

	ctype := resp.Header.Get("Content-Type")
	if ctype == "" {
		return "", fmt.Errorf("can't find Content-Type header")
	}

	return ctype, nil
}

func main() {
	ctype, err := contentType("https://linkedin.com")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(ctype)
	}
}
