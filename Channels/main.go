package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		// "http://google.com",
		"http://baidu.com",
		"http://bilibili.com/",
		// "http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func() {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}()
	}

	// for _, link := range links {
	// 	fmt.Println("checking link: ", link)
	// 	go checkLink(link, c)
	// }

	// fmt.Println(<-c)

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "   -> Link is down. Error: ", err)
		c <- link
		return
	}

	c <- link
	fmt.Println(link, "   -> Link is up")
}
