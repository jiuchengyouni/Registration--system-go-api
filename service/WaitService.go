package service

import (
	entiyPo "Registration-system/entiy/po"

	"github.com/jinzhu/gorm"
)

func GetWaiter(db *gorm.DB,department string,stuName string)int{
	var number int
	var wait []entiyPo.Wait
	db.Where("department=?",department).Find(&wait)
	for i:=0;i<len(wait);i++{
		if stuName==wait[i].StuName{
			number=i
			break
		}
	}
	return number
}