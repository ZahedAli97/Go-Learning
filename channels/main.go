package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	//for i := 0; i < len(links); i++ {
	//	//fmt.Println(<-c) // Because we want to wait for messages from every go routine - so basic for loop
	//}

	// infinite loop : for {}

	for l := range c { // run every single time, channel emits a value
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down!")
		c <- link
		return
	}
	fmt.Println(link, " is up!")
	c <- link
}
