package service

import (
	"errors"
	"tiktop/entity"
	"tiktop/global"
)

// 根据视频id列表查询评论数量列表
func QueryCommentCountListByVideoIdList(videoIdList *[]int64) (commentCountList []entity.VideoCommentCnt, err error) {
	result := global.DB.Model(&entity.Comment{}).Select("video_id", "count(video_id) as comment_cnt").Where("video_id in ?", *videoIdList).Group("video_id").Find(&commentCountList)
	if result.Error != nil {
		err = errors.New("likesList query failed")
		return
	}
	return
}
