package tmpl

//go:generate mockgen -destinition=mock_$GOFILE -package=$GOPACKAGE

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type User struct {
	ID            int64
	Name          string
	LastLoginDate time.Time
}

type UserPK struct {
	ID int64
}

func (u *User) toUserPK() UserPK {
	return UserPK{ID: u.ID}
}

type Users []*User

func (hs *Users) ToMap() map[UserPK]*User {
	m := make(map[UserPK]*User, len(*hs))
	for _, h := range *hs {
		m[h.toUserPK()] = h
	}
	return m
}

type UserModelRepository struct {
	client *sqlx.DB
}

func NewUserModelRepository(c *sqlx.DB) *UserModelRepository {
	return &UserModelRepository{
		client: c,
	}
}

func (r *UserModelRepository) Get(hk UserPK) (*User, error) {
	h := new(User)
	if err := r.client.Select(&h, "SELECT * FROM User WHERE id=?", hk.ID); err != nil {
		return nil, err
	}
	return h, nil
}

func (r *UserModelRepository) FindByName(name string) (Users, error) {
	var hs Users
	if err := r.client.Select(&hs, "SELECT * FROM User WHERE name=?", name); err != nil {
		return nil, err
	}
	return hs, nil
}
