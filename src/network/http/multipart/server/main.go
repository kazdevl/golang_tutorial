package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(_ http.ResponseWriter, r *http.Request) {
		fmt.Printf("Content-Type: %s\n", r.Header.Get("Content-Type"))
		defer r.Body.Close()
		defaultMaxMemory := 32 << 20
		if err := r.ParseMultipartForm(int64(defaultMaxMemory)); err != nil {
			log.Println(err)
		}

		if r.MultipartForm != nil {
			fmt.Println("Exist MultiPlatform")
			if r.MultipartForm.File != nil {
				fmt.Println("Exist MultiPlatform.File")
			}
			if r.MultipartForm.Value != nil {
				fmt.Println("Exist MultiPlatform.Value")
			}
		}
		fmt.Println("MultiPlatform.File")
		for key, datas := range r.MultipartForm.File {
			fmt.Printf("key: %s, data len: %d\n", key, len(datas))
		}
		fmt.Println("MultiPlatform.Value")
		for key, datas := range r.MultipartForm.Value {
			fmt.Printf("key: %s, data len: %d", key, len(datas))
			for _, str := range datas {
				fmt.Printf("str: %s", str)
			}
		}
		fmt.Println("Create Images from MultiPlatform.File")
		for _, images := range r.MultipartForm.File {
			for _, image := range images {
				fmt.Printf("create %s.png\n", image.Filename)
				f, err := os.Create(fmt.Sprintf("./images/%s.png", image.Filename))
				if err != nil {
					log.Println(err)
				}
				defer f.Close()
				content, err := image.Open()
				if err != nil {
					log.Println(err)
				}
				io.Copy(f, content)
			}
		}
	})
	http.ListenAndServe(":8080", nil)
}
