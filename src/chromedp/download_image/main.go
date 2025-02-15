package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewRemoteAllocator(context.Background(), "http://127.0.0.1:9222")
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx, chromedp.WithLogf(log.Printf))
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	done := make(chan bool)

	urlStr := "https://avatars.githubusercontent.com/u/33149672"

	var requestId network.RequestID
	chromedp.ListenTarget(ctx, func(v any) {
		switch ev := v.(type) {
		case *network.EventRequestWillBeSent:
			log.Printf("request will be sent: %v: %s", ev.RequestID, ev.Request.URL)
			if ev.Request.URL == urlStr {
				requestId = ev.RequestID
			}

		case *network.EventLoadingFinished:
			log.Printf("loading finished: %v", ev.RequestID)
			if ev.RequestID == requestId {
				close(done)
			}
		}
	})

	if err := chromedp.Run(ctx, chromedp.Navigate(urlStr)); err != nil {
		log.Fatal(err)
	}

	<-done
	var buf []byte
	if err := chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		var err error
		buf, err = network.GetResponseBody(requestId).Do(ctx)
		return err
	})); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("avatar.png", buf, 0644); err != nil {
		log.Fatal(err)
	}
	log.Printf("downloaded avatar.png")
}
