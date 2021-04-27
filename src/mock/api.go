package mock

import "fmt"

type API interface {
	Create(name string, age int) error
	Find(id int) (*User, error)
}

type User struct {
	ID   int
	Name string
	Age  int
}

type UserAPI struct {
	url string
}

func NewUserAPI(url string) API {
	return &UserAPI{url: url}
}

func (u *UserAPI) Create(name string, age int) error {
	// 本来であれば何らかの処理
	fmt.Printf("name: %v, age: %d\n", name, age)
	return nil
}

func (u *UserAPI) Find(id int) (*User, error) {
	if id == 1 {
		return &User{Name: "owner", Age: 100}, nil
	}

	return &User{Name: "mob", Age: 20}, nil
}

func wantExport() string {
	return "exported!"
}
