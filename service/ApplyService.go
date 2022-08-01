package service

import (
	entiy "Registration-system/entiy/po"

	"github.com/jinzhu/gorm"
)

func IsStuNum(db *gorm.DB,stuNum int)bool{
	var apply entiy.Apply
	db.Where("stu_num=?",stuNum).First(&apply)
	if apply.ID!=0{
		return true
	}
	return false
}