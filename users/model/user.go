package model

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	ID          string    `json:"id" gorm:"primaryKey"`
	Username    string    `json:"username" gorm:"not null"`
	FirstName   string    `json:"first_name" gorm:"not null"`
	LastName    string    `json:"last_name" gorm:"not null"`
	Email       string    `json:"email" gorm:"not null"`
	DateOfBirth string    `json:"date_of_birth" gorm:"not null"`
	Address     []Address `json:"address" gorm:"embedded"`
}
