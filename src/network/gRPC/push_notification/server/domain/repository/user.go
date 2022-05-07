package repository

type IFUserRepository interface {
	Register(string) error
	GetNames() ([]string, error)
}
