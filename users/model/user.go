package model

type User struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Username    string    `json:"username" gorm:"not null"`
	FirstName   string    `json:"first_name" gorm:"not null" db:"first_name"`
	LastName    string    `json:"last_name" gorm:"not null" db:"last_name"`
	Email       string    `json:"email" gorm:"not null"`
	DateOfBirth string    `json:"date_of_birth" gorm:"column:dob" db:"dob"`
	Address     []Address `json:"addresses" db:"-"`
}
