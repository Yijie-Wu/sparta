package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	NT     string `gorm:"size:64;not null"`
	Name   string `gorm:"size:128;not null"`
	Email  string `gorm:"size:256;not null"`
	Avatar string `gorm:"size:256;not null"`
}
