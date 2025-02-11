package main

import (
	"context"
	"log"
	"os"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
)

func main() {
	ctx, cancel := chromedp.NewRemoteAllocator(context.Background(), "ws://localhost:9222/")
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	var b1, b2 []byte
	if err := chromedp.Run(ctx,
		chromedp.Emulate(device.IPhone7landscape),
		chromedp.Navigate("https://www.whatsmyua.info/"),
		chromedp.CaptureScreenshot(&b1),
		chromedp.Emulate(device.Reset),
		chromedp.EmulateViewport(1920, 2000),
		chromedp.Navigate("https://www.whatsmyua.info/?a"),
		chromedp.CaptureScreenshot(&b2),
	); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("screenshot1.png", b1, 0o644); err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile("screenshot2.png", b2, 0o644); err != nil {
		log.Fatal(err)
	}

	log.Printf("screenshot1 saved")
	log.Printf("screenshot2 saved")
}
