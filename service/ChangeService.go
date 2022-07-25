package service

import (
	"Registration-system/model"

	"github.com/jinzhu/gorm"
)

func Waitline(db *gorm.DB,department string)[]model.WaitInfo {
	var waitInfo []model.WaitInfo
	db.Where("department=? ",department).Find(&waitInfo)
	return waitInfo
}