package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/friendsofgo/errors"
	"golang.org/x/sync/errgroup"
)

var targetURL = "http://localhost:8888"
var targetFile = "sample.txt"

func main() {
	// 部分的なリクエストに対応しているのかの確認
	contentLength, err := getTargetContentLength(targetFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Content-length: ", contentLength)

	// 5並列でリクエストを投げてデータを保存する
	path, err := filepath.Abs("./")
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Join(path, "download")
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatal(err)
	}

	dls := createDonwloaders(dir, targetFile, 5, contentLength)
	if err := dls.ParallelDownload(context.Background()); err != nil {
		log.Fatal(err)
	}

	if err := bindFiles(dir, targetFile, 5); err != nil {
		log.Fatal(err)
	}
}

func getTargetContentLength(target string) (int64, error) {
	req, _ := http.NewRequest(http.MethodHead, targetURL+"/"+target, nil)
	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	if resp.Header.Get("Accept-Ranges") != "bytes" {
		return 0, errors.WithStack(errors.New("Doee not support range request"))
	}
	return resp.ContentLength, nil
}

type Downloader struct {
	Client   *http.Client
	FileName string
	Request  *http.Request
}

func createDonwloaders(dir, target string, procs int, contentLength int64) Downloaders {
	dls := make(Downloaders, 0, procs)
	for i := 0; i < procs; i++ {
		r := newRange(i, procs, contentLength)
		req, _ := http.NewRequest(http.MethodGet, targetURL+"/"+target, nil)
		req.Header.Set("Range", r.createHeaderValue())
		dls = append(dls, Downloader{
			Client:   &http.Client{},
			FileName: filepath.Join(dir, fmt.Sprintf("%d_%s", i, target)),
			Request:  req,
		})
	}
	return dls
}

type Downloaders []Downloader

func (ds Downloaders) ParallelDownload(ctx context.Context) error {
	eg, _ := errgroup.WithContext(ctx)
	for _, d := range ds {
		d := d
		eg.Go(func() error {
			fmt.Printf("dl request header: %s\n", d.Request.Header.Get("Range"))
			resp, err := d.Client.Do(d.Request)
			if err != nil {
				return errors.WithStack(err)
			}
			defer resp.Body.Close()

			f, err := os.Create(d.FileName)
			if err != nil {
				return errors.WithStack(err)
			}
			defer f.Close()

			if _, err := io.Copy(f, resp.Body); err != nil {
				return errors.WithStack(err)
			}
			return nil
		})
	}
	return eg.Wait()
}

type Range struct {
	Start int64
	End   int64
}

func newRange(index, procs int, contentLength int64) Range {
	rangeSize := contentLength / int64(procs)
	start := rangeSize * int64(index)
	if index == procs-1 {
		return Range{
			Start: start,
			End:   contentLength,
		}
	}
	return Range{
		Start: start,
		End:   start + rangeSize - 1,
	}
}

func (r Range) createHeaderValue() string {
	fmt.Printf("bytes=%d-%d\n", r.Start, r.End)
	return fmt.Sprintf("bytes=%d-%d", r.Start, r.End)
}

func bindFiles(dir, target string, procs int) error {
	f, err := os.Create(filepath.Join(dir, target))
	if err != nil {
		return err
	}
	defer f.Close()

	for i := 0; i < procs; i++ {
		partF, err := os.Open(filepath.Join(dir, fmt.Sprintf("%d_%s", i, target)))
		if err != nil {
			return err
		}
		defer partF.Close()

		if _, err := io.Copy(f, partF); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}
