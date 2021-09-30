package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://127.0.0.1",
		"http://google.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for ii, link := range links {
		fmt.Println(ii)
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		fmt.Println(err)
		return
	}

	fmt.Println(link, "is up!")
}
