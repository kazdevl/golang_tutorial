package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

// NOTE: 事前にdocker run -d -p 9222:9222 --rm --name headless-shell chromedp/headless-shellでheadless-shellを起動しておく
// NOTE: 実行中のheadless-shellを利用するには、NewRemoteAllocatorを使う
func main() {
	ctx, cancel := chromedp.NewRemoteAllocator(context.Background(), "ws://localhost:9222/")
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var example string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://pkg.go.dev/time`),
		chromedp.WaitVisible(`body > footer`),
		chromedp.Click(`#example-After`, chromedp.NodeVisible),
		chromedp.Value(`#example-After textarea`, &example),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Go's time.After: \n%s", example)
}
