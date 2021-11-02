package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"log"
)

func main() {
	worklist := make(chan []string)

	go func() { worklist <- []string{"https://www.baidu.com", "https://chenxf.org/"} }()

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

var  tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	if err != nil {
		log.Println(err)
	}
	<-tokens
	return list
}
