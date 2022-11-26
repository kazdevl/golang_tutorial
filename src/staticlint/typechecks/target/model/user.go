package model

type User struct {
	ID   int64
	Name string
}

type UserPK struct {
	ID int64
}

func (u *User) toUserPK() UserPK {
	return UserPK{
		ID: u.ID,
	}
}

type Users []*User

func (us *Users) ToMap() map[UserPK]*User {
	m := make(map[UserPK]*User, len(*us))
	for _, u := range *us {
		m[u.toUserPK()] = u
	}
	return m
}
