package model

import "time"

type Post struct {
	ID        int64
	UserID    int64
	Content   string
	CreatedAt time.Time
}

type PostPK struct {
	ID int64
}

func (p *Post) toPostPK() PostPK {
	return PostPK{
		ID: p.ID,
	}
}

type Posts []*Post

func (ps *Posts) ToMap() map[PostPK]*Post {
	m := make(map[PostPK]*Post, len(*ps))
	for _, p := range *ps {
		m[p.toPostPK()] = p
	}
	return m
}
