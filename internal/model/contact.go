package model

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Name    string `json:"name"`
	Email   string `json:"email" gorm:"unique"`
	PhoneNo string `json:"phone_no" gorm:"unique"`
	Address string `json:"address"`
}
