package repository

import (
	"app/network/gRPC/push_notification/server/domain/entity"
)

type IFTweetRepository interface {
	Post(string) error
	GetByIDs([]string) (entity.Tweet, error)
}
