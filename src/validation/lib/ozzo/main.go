package main

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func init() {
	validation.ErrRequired = validation.ErrRequired.SetMessage("必ず値を入力してください")
}

func main() {
	data := ""
	err := validation.Validate(data, validation.Required, validation.Length(5, 10), is.URL)
	fmt.Println(err)
	fmt.Println("*********")
	err = validation.Validate(data, validation.By(customValidation))
	fmt.Println(err)
}

func customValidation(value any) error {
	s, _ := value.(string)
	if s != "hello" {
		return fmt.Errorf("value must be 'hello'")
	}
	return nil
}
