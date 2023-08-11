package service

import (
	"errors"
	"tiktop/entity"
	"tiktop/global"
)

func QueryLikesListByVideoIdList(videoIdList *[]int64) (likesList []entity.VideoLikeCnt, err error) {
	//获取每个视频的点赞数量
	var videoLikeCount []entity.VideoLikeCnt
	result := global.DB.Model(&entity.Like{}).Select("video_id", "count(video_id) as like_cnt").Where("video_id in ?", *videoIdList).Group("video_id").Find(&videoLikeCount)
	if result.Error != nil {
		err = errors.New("likesList query failed")
		return
	}
	return
}
