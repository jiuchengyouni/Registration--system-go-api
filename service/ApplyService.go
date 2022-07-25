package service

import (
	"Registration-system/model"

	"github.com/jinzhu/gorm"
)

func IsStuNum(db *gorm.DB,stuNum int)bool{
	var applyInfo model.ApplyInfo
	db.Where("stu_num=?",stuNum).First(&applyInfo)
	if applyInfo.ID!=0{
		return true
	}
	return false
}