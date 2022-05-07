package entity

type User struct {
	Name string
}

func NewUser(input string) *User {
	if len(input) == 0 {
		return nil
	}
	return &User{Name: input}
}
