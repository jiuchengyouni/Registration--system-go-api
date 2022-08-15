package service

import "github.com/jinzhu/gorm"

func IsStuNum_login(db *gorm.DB, stuNum string) bool {
	return IsStuNum(db, stuNum)
}
