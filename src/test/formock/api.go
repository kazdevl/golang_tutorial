package formock

import "fmt"

type FormValidater interface {
	Validate(string) (string, error)
}

type API struct {
	Validater FormValidater
}

func (a *API) RegisterData(input string) error {
	data, err := a.Validater.Validate(input)
	if err != nil {
		return err
	}

	fmt.Printf("input that is registerd into database: %v", data)
	// dataをdbに登録する処理
	return nil
}
