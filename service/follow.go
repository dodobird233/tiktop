package service

import (
	"tiktop/entity"
	"tiktop/global"
)

func QueryFollowOrNot(userId int64, toId int64) (res bool, err error) {
	var follow []entity.Follow
	result := global.DB.Model(&entity.Follow{}).Where("user_id = ? and follow_userid = ?", userId, toId).Find(&follow)
	if result.RowsAffected == 0 {
		return false, nil
	}
	if result.Error != nil {
		err = result.Error
		return
	}
	return true, nil
}
