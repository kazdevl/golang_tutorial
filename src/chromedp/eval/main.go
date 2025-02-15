package main

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewRemoteAllocator(context.Background(), "http://127.0.0.1:9222")
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	var res []string
	err := chromedp.Run(
		ctx,
		chromedp.Navigate("https://www.google.com"),
		chromedp.EvaluateAsDevTools(`
		Object.keys(window);
	`, &res))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("res: %v", res)
}
