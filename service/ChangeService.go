package service

import (
	entiyParm "Registration-system/entiy/parm"
	entiyPo "Registration-system/entiy/po"
	"fmt"

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

func Reset(db *gorm.DB,waitInfo entiyParm.WaitInfo)bool{
	wait:=entiyPo.Wait{}
	db.Where("department=? AND stu_num=?",waitInfo.Department,waitInfo.StuNum).First(&wait)
	if wait.ID==0{
		return false
	}
	db.Exec("update waits set state = NULL where stu_num=?",waitInfo.StuNum)
	fmt.Println("更新了",db.RowsAffected,"条数据")
	///db.Model(&wait).Where("department=? AND stu_num=?",waitInfo.Department,waitInfo.StuNum).Update(map[string]interface{}{"state":nil})
	return true
}

func Next(db *gorm.DB,department string)bool{
	wait:=entiyPo.Wait{}
	db.Where(map[string]interface{}{"department":department, "state":nil}).First(&wait)
	db.Model(&wait).Where(map[string]interface{}{"department":department, "state":nil}).Update(map[string]interface{}{"state":"finished"})
	return true
}