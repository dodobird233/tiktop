package global

import (
	"gorm.io/gorm"
)

var ( // 系统配置信息
	DB *gorm.DB // 数据库接口
)
