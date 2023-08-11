package entity

type Video struct {
	VideoId  int64  `gorm:"column:video_id;primary_key;NOT NULL"`
	VideoUrl string `gorm:"column:video_url;type:varchar(500)"`
	CoverUrl string `gorm:"column:cover_url;type:varchar(500)"`
	Title    string `gorm:"column:title;type:varchar(100)"`
	UserId   int64  `gorm:"column:user_id;NOT NULL"`
}

type VideoLikeCnt struct {
	VideoId int64 `gorm:"column:video_id;primary_key;NOT NULL"`
	LikeCnt int64 `gorm:"column:like_cnt;NOT NULL"`
}
