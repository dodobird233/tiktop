package service

import (
	"tiktop/entity"
	"tiktop/global"
)

func QueryVideoListByUserId(userId int64) (videoIdList []int64, err error) {
	var videoList []entity.Video
	result := global.DB.Where("user_id = ?", userId).Find(&videoList)
	if result.Error != nil {
		err = result.Error
		return
	}
	if result.RowsAffected == 0 {
		return
	}
	//构造视频id表
	videoCnt := result.RowsAffected
	videoIdList = make([]int64, videoCnt)
	for i, video := range videoList {
		videoIdList[videoCnt-int64(i)-1] = video.VideoId
	}
	return
}
