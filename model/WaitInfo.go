package model

import "github.com/jinzhu/gorm"

type WaitInfo struct {
	gorm.Model
	Department string `gorm:"type:varchar(20);not null"`
	StuName         string `gorm:"type:varchar(20);not null"`
}