package model

import "github.com/jinzhu/gorm"

type ChangeInfo struct{
	gorm.Model
	Department string `gorm:"type:varchar(20);not null"`
	StuNum int `gorm:"type:varchar(20);not null"`
}