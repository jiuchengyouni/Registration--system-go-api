package service

import (
	entiyPo "Registration-system/entiy/po"

	"github.com/jinzhu/gorm"
)

func GetWaiter(db *gorm.DB,department string,stuNum string)int{
	var number int
	var wait []entiyPo.Wait
	db.Where(map[string]interface{}{"department":department, "state":nil}).Find(&wait)
	for i:=0;i<len(wait);i++{
		if stuNum==wait[i].StuNum{
			number=i
			break
		}
	}
	return number
}