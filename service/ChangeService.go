package service

import (
	entiyPo "Registration-system/entiy/po"
	"github.com/jinzhu/gorm"
)

func Waitline(db *gorm.DB,department string)[]entiyPo.Wait {
	var wait []entiyPo.Wait
	db.Where("department=? ",department).Find(&wait)
	return wait
}

func Advance(db *gorm.DB,wait2 entiyPo.Wait)bool{
	var waitList[] entiyPo.Wait
	db.Where("department=? AND id<?",wait2.Department,wait2.ID).Find(&waitList)
	wait:=waitList[len(waitList)-1]
	wait2.StuNum,wait.StuNum=wait.StuNum,wait2.StuNum
	wait2.StuName,wait.StuName=wait.StuName,wait2.StuName
	wait2.PhoneNum,wait.PhoneNum=wait.PhoneNum,wait2.PhoneNum
	db.Save(&wait)
	db.Save(&wait2)
	return true
}

func Delay(db *gorm.DB,wait entiyPo.Wait)bool{
	var wait2 entiyPo.Wait
	db.Where("department=? AND id>?",wait.Department,wait.ID).First(&wait2)
	wait2.StuNum,wait.StuNum=wait.StuNum,wait2.StuNum
	wait2.StuName,wait.StuName=wait.StuName,wait2.StuName
	wait2.PhoneNum,wait.PhoneNum=wait.PhoneNum,wait2.PhoneNum
	db.Save(&wait)
	db.Save(&wait2)
	return true
}