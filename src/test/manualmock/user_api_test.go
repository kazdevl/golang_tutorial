package mock_test

import (
	"errors"
	"testing"

	mock "github.com/kazdevl/golang_tutorial/test/manualmock"
)

type MockSQL struct{}

// func (m *MockSQL) CreateUser(name string) error {
// 	fmt.Printf("name: %v", name)
// 	return nil
// }

func (m *MockSQL) FindUser(id int) (*mock.User, error) {
	if id <= 0 {
		return &mock.User{}, errors.New("No such id exists.")
	}
	return &mock.User{ID: id, Name: "Hello World"}, nil
}

func TestUserAPI(t *testing.T) {
	target := &mock.UserAPI{Handler: &MockSQL{}}
	expected := "Hello World"
	_, name, _ := target.Get(1)
	t.Logf("expected: %v, resut: %v", expected, name)
	if _, _, err := target.Get(-100); err != nil {
		t.Error(err)
	}
}
