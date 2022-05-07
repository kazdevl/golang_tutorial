package entity

import "time"

type Tweet struct {
	ID        int64
	Message   string
	CreatedAt time.Time
}
