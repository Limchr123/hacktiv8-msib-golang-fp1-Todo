package entity

import "gorm.io/gorm"

//Todo represents the model of todo
type Todo struct {
	gorm.Model
	Title     string `gorm:"not null"`
	Completed bool   `gorm:"not null"`
}
