package main

import (
	"log"

	"gopkg.in/go-playground/validator.v9"
)

func main() {
	err := validator.New()
	log.Println(err)
}
