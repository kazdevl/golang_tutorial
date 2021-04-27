package mock

import "fmt"

type UserAPI struct {
	Handler SqlHandlerInterface
}

func (u_api *UserAPI) Get(id int) (int, string, error) {
	user, err := u_api.Handler.FindUser(id)
	if err != nil {
		return 0, "null", err
	}
	fmt.Printf("user: %+v\n", user)
	return user.ID, user.Name, nil
}

// func (u_api *UserAPI) Create(firstname, lastname string) error {
// 	err := u_api.Handler.CreateUser(firstname + " " + lastname)
// 	return err
// }
