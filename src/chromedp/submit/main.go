package main

import (
	"context"
	"log"
	"strings"

	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewRemoteAllocator(context.Background(), "http://127.0.0.1:9222")
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx, chromedp.WithDebugf(log.Printf))
	defer cancel()

	var res string
	if err := chromedp.Run(ctx, submit("https://github.com/search", `//input[@name="q"]`, `chromedp`, &res)); err != nil {
		log.Fatal(err)
	}

	log.Printf("got: %s", strings.TrimSpace(res))
}

func submit(urlStr, sel, q string, res *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlStr),
		chromedp.WaitVisible(sel),
		chromedp.SendKeys(sel, q),
		chromedp.Submit(sel),
		chromedp.WaitVisible(`//*[contains(., 'respository results')]`),
		chromedp.Text(`(//*//ul[contains(@class, "repo-list")]/li[1]//p)[1]`, res),
	}
}
