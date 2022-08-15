package entiy

import "github.com/jinzhu/gorm"

type Apply struct {
	gorm.Model
	Department1 string `gorm:"type:varchar(20);not null"`
	Department2 string `gorm:"type:varchar(20);not null"`
	PhoneNum string `gorm:"type:varchar(20);not null"`
	StuName string `gorm:"type:varchar(20);not null"`
	StuNum string `gorm:"type:varchar(20);not null"`
}