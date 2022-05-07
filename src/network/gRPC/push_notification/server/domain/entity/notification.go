package entity

type Notification struct {
	ID          int64
	UserID      int64
	TweetID     int64
	IsMensioned bool
}
