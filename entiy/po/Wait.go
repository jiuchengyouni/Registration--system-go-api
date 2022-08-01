package entiy

import "github.com/jinzhu/gorm"

type Wait struct {
	gorm.Model
	Department string `gorm:"type:varchar(20);not null"`
	PhoneNum   int    `gorm:"type:int(20);not null"`
	StuName    string `gorm:"type:varchar(20);not null"`
	StuNum     int    `gorm:"type:int(20);not null"`
}