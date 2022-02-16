package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://netflix.com",
	}
	
	var wg sync.WaitGroup
	
	wg.Add(len(links))

	for _, link := range links {
		go func (l string) {
			checkLinkStatus(l)
			
			wg.Done()
		}(link)
	}
	
	wg.Wait()
}

func checkLinkStatus(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "is down")

		return
	}

	fmt.Println(link, "is up!")
}
