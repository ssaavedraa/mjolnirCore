package models

import (
	"time"

	"gorm.io/gorm"
)

type Shift struct {
	gorm.Model
	ClockIn  *time.Time `gorm:"default:current_timestamp"`
	ClockOut *time.Time `gorm:"default:null"`
	Start    *time.Time `gorm:"default:null"`
	End      *time.Time `gorm:"default:null"`
	UserID   uint       `json:"userId"`
}
