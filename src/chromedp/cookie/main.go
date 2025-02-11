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
					return fmt.Errorf("クッキーの設定に失敗しました: %w", err)
				}
			}
			return nil
		}),
		chromedp.Navigate(host),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// 10秒のタイムアウトを設定
			ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()

			// #result要素が存在するかどうかを確認
			var exists bool
			if err := chromedp.Run(ctx, chromedp.QueryAfter("#result", func(ctx context.Context, n ...*cdp.Node) error {
				exists = len(n) > 0
				return nil
			})); err != nil {
				return fmt.Errorf("ページの読み込みに失敗しました: %w", err)
			}

			if !exists {
				*res = "" // 要素が存在しない場合は空文字を設定
				return nil
			}

			// 要素が存在する場合のみテキストを取得
			if err := chromedp.Text("#result", res, chromedp.ByID, chromedp.NodeVisible).Do(ctx); err != nil {
				return fmt.Errorf("テキストの取得に失敗しました: %w", err)
			}
			return nil
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			cookies, err := storage.GetCookies().Do(ctx)
			if err != nil {
				return fmt.Errorf("クッキーの取得に失敗しました: %w", err)
			}
			for i, cookie := range cookies {
				log.Printf("chrome cookie %d: %s\n", i, cookie.Name)
			}
			return nil
		}),
	}
}
