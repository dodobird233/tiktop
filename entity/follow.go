package entity

type Follow struct {
	UserId       int64 `gorm:"column:user_id;primary_key;NOT NULL"`
	FollowUserId int64 `gorm:"column:follow_userid;NOT NULL"`
}
