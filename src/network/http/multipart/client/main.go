package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// type contentHasClose interface {
// 	Close() error
// }

// type closer struct {
// 	contents []contentHasClose
// }

// func (c *closer) close() error {
// 	var err error = nil
// 	for _, content := range c.contents {
// 		err = content.Close()
// 	}
// 	return err
// }

func main() {
	paths := getImagesPaths("./asset")
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for i, path := range paths {
		image, _ := os.Open(path)
		defer image.Close()
		part, _ := writer.CreateFormFile("image", fmt.Sprintf("sample%d", i))
		io.Copy(part, image)
	}
	writer.Close()
	r, _ := http.NewRequest("POST", "http://localhost:8080", body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	c := &http.Client{}
	c.Do(r)
}

func getImagesPaths(dir string) []string {
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	imagesPaths := make([]string, len(fileInfos))
	for i, fileinfo := range fileInfos {
		imagesPaths[i] = filepath.Join(dir, fileinfo.Name())
	}
	return imagesPaths
}
