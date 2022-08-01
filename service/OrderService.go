package service

import (
	entiyPo "Registration-system/entiy/po"

	"github.com/jinzhu/gorm"
)

func IsOrder(db *gorm.DB,stuNum int)entiyPo.Apply{
	var apply entiyPo.Apply
	db.Where("stu_num=?",stuNum).First(&apply)
	return apply
}