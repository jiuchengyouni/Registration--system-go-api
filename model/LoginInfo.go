package model

import "github.com/jinzhu/gorm"

type LoginInfo struct {
	gorm.Model
	PhoneNum int `gorm:"type:int(20);not null"`
	StuName string `gorm:"type:varchar(20);not null"`
	StuNum int `gorm:"type:int(20);not null"`
}