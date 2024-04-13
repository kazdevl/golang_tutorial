package main

import (
	"fmt"

	"github.com/go-faker/faker/v4"
)

type LangEtc struct {
	Eng string `faker:"lang=eng"`
	Kor string `faker:"lang=kor"`
	Jpn string `faker:"lang=jpn"`
	Chi string `faker:"lang=chi"`
}

func createDummyLang() *LangEtc {
	v := LangEtc{}
	err := faker.FakeData(&v)
	if err != nil {
		fmt.Println(err)
	}
	return &v
}
