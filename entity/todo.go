package entity

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title     string `gorm:"not null"`
	Completed bool   `gorm:"not null"`
}
