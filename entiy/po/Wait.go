package entiy

import "github.com/jinzhu/gorm"

type Wait struct {
	gorm.Model
	Department string 	`gorm:"type:varchar(20);not null"`
	PhoneNum    string    `gorm:"type:varchar(20);not null"`
	StuName    string 	`gorm:"type:varchar(20);not null"`
	StuNum     string    `gorm:"type:varchar(20);not null"`
	State      string	`gorm:"type:varchar(20);"`
}