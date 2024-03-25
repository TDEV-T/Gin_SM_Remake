package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username       string `gorm:"unique;not null" json:"username"`
	Password       string `gorm:"not null" json:"password"`
	DisplayName    string `json:"displayname"`
	Email          string `gorm:"unique:not null" json:"email"`
	PrivateAccount string `gorm:"default:false" json:"privateacc"`
}
