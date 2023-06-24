package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	NT       string `gorm:"size:64;not null" json:"nt"`
	Name     string `gorm:"size:128;not null" json:"name"`
	Email    string `gorm:"size:256;not null" json:"email"`
	Avatar   string `gorm:"size:256" json:"avatar"`
	Password string `gorm:"size:256;not null" json:"-"`
}
