package models

import "gorm.io/gorm"

type Example struct {
	gorm.Model // Adds CreatedAt, ID, etc
	Name       string
	Value      string
}
