package service

import "github.com/jinzhu/gorm"

func IsStuNum_login(db *gorm.DB, stuNum int) bool {
	return IsStuNum(db,stuNum)
}