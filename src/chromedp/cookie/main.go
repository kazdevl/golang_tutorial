package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/storage"
	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewRemoteAllocator(context.Background(), "ws://localhost:9222/")
	defer cancel()
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	var res string
	if err := chromedp.Run(ctx, setCookies(
		fmt.Sprintf("http://localhost:%d/cookie/result", 8544), &res,
		"cookie1", "value1",
		"cookie2", "value2",
	)); err != nil {
		log.Fatal(err)
	}

	log.Printf("chrome received cookies: %s", res)
}

func setCookies(host string, res *string, cookies ...string) chromedp.Tasks {
	if len(cookies)%2 != 0 {
		panic("length of cookies must be divisible by 2")
	}
	return chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
			for i := 0; i < len(cookies); i += 2 {
				if err := network.SetCookie(cookies[i], cookies[i+1]).WithExpires(&expr).WithDomain("localhost").WithHTTPOnly(true).Do(ctx); err != nil {
					return err
				}
			}
			return nil
		}),
		chromedp.Navigate(host),
		chromedp.Text(`#result`, res, chromedp.ByID, chromedp.NodeVisible),
		chromedp.ActionFunc(func(ctx context.Context) error {
			cookies, err := storage.GetCookies().Do(ctx)
			if err != nil {
				return err
			}
			for i, cookie := range cookies {
				log.Printf("chrome cookie %d: %s\n", i, cookie.Name)
			}

			return nil
		}),
	}
}
