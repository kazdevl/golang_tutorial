package main

import (
	"fmt"

	"github.com/go-faker/faker/v4"
)

type Name struct {
	Value string `faker:"name,unique"`
}

func createDummyUniques() []*Name {
	vs := make([]*Name, 0, 3)
	for range 3 {
		v := Name{}
		err := faker.FakeData(&v)
		if err != nil {
			fmt.Println(err)
		}
		vs = append(vs, &v)
	}
	faker.ResetUnique()
	return vs
}
