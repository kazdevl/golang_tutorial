package main

import (
	"fmt"
	"log"

	"github.com/mmcdole/gofeed"
)

func main() {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://zenn.dev/topics/go/feed")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(feed.Title)
	for _, item := range feed.Items {
		fmt.Println(item.Title)
		fmt.Println(item.Link)
		fmt.Println(item.Updated)
		fmt.Println("**********")
	}
}
