package model

import "gorm.io/gorm"

type Invoice struct {
	gorm.Model
	FileName string
	FileUrl string
	UserID uint
}