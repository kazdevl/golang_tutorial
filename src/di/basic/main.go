package main

import (
	"fmt"
	"log"

	"go.uber.org/dig"
)

type Repository interface {
	Create(id, name string) error
}

type SampleRepository struct {
	Client string
}

func NewSampleReository(c string) Repository {
	return &SampleRepository{c}
}

func (r *SampleRepository) Create(id, name string) error {
	fmt.Printf("call in sample repository create: id=%s, name=%s, client=%s\n", id, name, r.Client)
	return nil
}

type Usecase interface {
	Create(name string) error
}

type SampleUsecase struct {
	Repository Repository
}

func NewSampleUsecase(r Repository) Usecase {
	return &SampleUsecase{r}
}

func (u *SampleUsecase) Create(name string) error {
	fmt.Printf("call in sample usecase create: name=%s\n", name)
	u.Repository.Create("hogehoge", name)
	return nil
}

func main() {
	c := dig.New()
	c.Provide(func() string {
		return "sample"
	})
	c.Provide(NewSampleReository)
	c.Provide(NewSampleUsecase)
	if err := c.Invoke(func(usecase Usecase) error {
		return usecase.Create("sample")
	}); err != nil {
		log.Fatal(err)
	}
}
