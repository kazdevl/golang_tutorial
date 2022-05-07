package repository

import (
	"app/network/gRPC/push_notification/server/domain/entity"
)

type IFNotificationRepository interface {
	Add(*entity.Notification) error
	GetByUserID(id int64) ([]*entity.Notification, error)
	DeleteByIDs([]int64) error
}
