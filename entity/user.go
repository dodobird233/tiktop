package entity

type UserData struct {
	UserId          int64
	Name            string
	FollowCount     int64
	FollowerCount   int64
	IsFollow        bool
	Avatar          string
	BackgroundImage string
	Signature       string
	TotalFavorited  int64
	WorkCount       int64
	FavoriteCount   int64
}

type User struct {
	UserId          int64  `gorm:"column:user_id;primary_key;NOT NULL"`
	UserName        string `gorm:"column:user_name;type:varchar(100)"`
	Password        string `gorm:"column:password;type:varchar(100)"`
	Avatar          string `gorm:"column:avatar;type:varchar(100)"`
	BackgroundImage string `gorm:"column:background_image;type:varchar(100)"`
	Signature       string `gorm:"column:signature;type:varchar(200)"`
}
