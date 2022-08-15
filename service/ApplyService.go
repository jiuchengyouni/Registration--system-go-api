package service

import (
	entiyPo "Registration-system/entiy/po"
	
	"github.com/jinzhu/gorm"
)

func IsStuNum(db *gorm.DB,stuNum string)bool{
	var apply entiyPo.Apply
	db.Where("stu_num=?",stuNum).First(&apply)
	if apply.ID!=0{
		return true
	}
	return false
}

func GetApply(db *gorm.DB)[]entiyPo.Apply{
	var applyList[] entiyPo.Apply
	db.Find(&applyList)
	return applyList
}