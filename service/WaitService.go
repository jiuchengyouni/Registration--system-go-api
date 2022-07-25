package service

import (
	"Registration-system/model"

	"github.com/jinzhu/gorm"
)

func GetWaiter(db *gorm.DB,department string,stuName string)int{
	var number int
	var waitInfo []model.WaitInfo
	db.Where("department=?",department).Find(&waitInfo)
	for i:=0;i<len(waitInfo);i++{
		if stuName==waitInfo[i].StuName{
			number=i
			break
		}
	}
	return number
}