package model

import (
	"gorm.io/gorm"
	"sparta/utils"
)

type User struct {
	gorm.Model
	NT       string `gorm:"size:64;not null" json:"nt"`
	Name     string `gorm:"size:128;not null" json:"name"`
	Email    string `gorm:"size:256;not null" json:"email"`
	Avatar   string `gorm:"size:256" json:"avatar"`
	Password string `gorm:"size:256;not null" json:"-"`
}

func (m *User) Encrypt() error {
	hash, err := utils.Encrypt(m.Password)
	if err == nil {
		m.Password = hash
	}
	return err
}

func (m *User) BeforeCreate(orm *gorm.DB) error {
	return m.Encrypt()
}
