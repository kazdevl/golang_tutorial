package mock

import "database/sql"

// mockしたい機能
type User struct {
	ID   int
	Name string
}

type SqlHandlerInterface interface {
	// CreateUser(name string) error
	FindUser(id int) (*User, error)
}

type SqlHandler struct {
	DB *sql.DB
}

func NewSqlHandler(db *sql.DB) SqlHandlerInterface {
	return &SqlHandler{DB: db}
}

// func (sh *SqlHandler) CreateUser(name string) error {
// 	_, err := sh.DB.Exec("INSERT INTO USER(name) VALUES (?)", name)
// 	return err
// }

func (sh *SqlHandler) FindUser(id int) (*User, error) {
	var user *User
	err := sh.DB.QueryRow("SELECT * FROM USER WHERE id == ?", id).Scan(user)
	return user, err
}
