package service

import (
	"Registration-system/model"

	"github.com/jinzhu/gorm"
)

func IsOrder(db *gorm.DB,stuNum int)model.ApplyInfo{
	var applyInfo model.ApplyInfo
	db.Where("stu_num=?",stuNum).First(&applyInfo)
	return applyInfo
}