package model

import (
	"time"

	"gorm.io/gorm"
)

type Shift struct {
	gorm.Model
	ClockIn time.Time
	ClockOut time.Time
	UserID uint
}