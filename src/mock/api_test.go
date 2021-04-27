package mock_test

import (
	"app/mock"
	"errors"
	"testing"
)

type MockUserAPI struct{}

func NewMockUserAPI() mock.API {
	return &MockUserAPI{}
}
func (mu *MockUserAPI) Create(name string, age int) error {
	return errors.New("errorを発生させているMockです")
}

func (mu *MockUserAPI) Find(id int) (*mock.User, error) {
	return &mock.User{ID: id, Name: "hello", Age: 10}, nil
}

type APIRegister struct {
	Instance mock.API
}

func TestMock(t *testing.T) {
	api := &APIRegister{NewMockUserAPI()}
	if err := api.Instance.Create("Hello", 100); err != nil {
		t.Error(err.Error())
	}

	if user, err := api.Instance.Find(1); err != nil {
		t.Error("cannot find")
	} else {
		t.Logf("user: %+v", user)
	}
}

func TestRealAPI(t *testing.T) {
	api := &APIRegister{mock.NewUserAPI("hello")}
	if err := api.Instance.Create("Hello", 100); err != nil {
		t.Error(err.Error())
	}

	if user, err := api.Instance.Find(1); err != nil {
		t.Error("cannot find")
	} else {
		t.Logf("user: %+v", user)
	}
}
