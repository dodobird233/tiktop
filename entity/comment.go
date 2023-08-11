package entity

type Comment struct {
	CommentId int64  `gorm:"column:comment_id;primary_key;NOT NULL"`
	UserId    int64  `gorm:"column:user_id;NOT NULL"`
	VideoId   int64  `gorm:"column:video_id;NOT NULL"`
	Content   string `gorm:"column:content_id;NOT NULL;type:varchar(300)"`
}
