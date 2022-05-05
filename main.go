package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func (p Page) String() string {
	return fmt.Sprintf("Page: %v, Size: %v", p.URL, p.Size)
}

func main() {
	urls := []string{"https://example.com/", "https://golang.org/", "https://golang.org/doc"}
	c := make(chan Page, len(urls))
	for i := 0; i < len(urls); i++ {
		go responseSize(urls[i], c)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}

}

func responseSize(url string, c chan Page) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	c <- Page{url, len(body)}
}
