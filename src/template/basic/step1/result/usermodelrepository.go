package usermodelrepository

//go:generate mockgen -destinition=mock_$GOFILE -package=$GOPACKAGE

import (
    "time"
)

type User struct {
    ID int64
    Name string
    LastLoginDate time.Time
}

type UserPK struct {
    ID int64
}