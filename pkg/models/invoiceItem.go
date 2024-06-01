package models

import "gorm.io/gorm"

type InvoiceItem struct {
	gorm.Model
	Description string `json:"string"`
	Quantity float64
	UnitPrice float64
	TotalPrice float64
	InvoiceID uint
}