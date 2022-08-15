package service

import (
	entiyParm "Registration-system/entiy/parm"
	entiyPo "Registration-system/entiy/po"

	"github.com/jinzhu/gorm"
)

func IsOrder(db *gorm.DB,stuNum string)entiyPo.Apply{
	var apply entiyPo.Apply
	db.Where("stu_num=?",stuNum).First(&apply)
	return apply
}

func Order(db *gorm.DB,orderInfo entiyParm.OrderInfo)bool{
	var apply entiyPo.Apply
	db.Where("stu_num=?",orderInfo.StuNum).First(&apply)
	if apply.ID==0{
		return false
	}
	var wait[] entiyPo.Wait
	db.Where("stu_num=?",orderInfo.StuNum).Find(&wait)
	if len(wait)>0{
		wait[0].Department=orderInfo.Department1
		wait[1].Department=orderInfo.Department2
		db.Updates(wait)
	}else{		
		wait:=entiyPo.Wait{
			Department: orderInfo.Department1,
			StuName: apply.StuName,
			StuNum: orderInfo.StuNum,
			PhoneNum: apply.PhoneNum,
		}
		db.Create(&wait)
		wait=entiyPo.Wait{
			Department: orderInfo.Department2,
			StuName: apply.StuName,
			StuNum: orderInfo.StuNum,
			PhoneNum: apply.PhoneNum,
		}
		db.Create(&wait)
	}
	return true
}