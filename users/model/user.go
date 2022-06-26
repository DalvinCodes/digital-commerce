package model

type User struct {
	ID          string `json:"id" gorm:"primaryKey"`
	Username    string `json:"username" gorm:"not null"`
	FirstName   string `json:"first_name" gorm:"not null" db:"first_name"`
	LastName    string `json:"last_name" gorm:"not null" db:"last_name"`
	Email       string `json:"email" gorm:"not null"`
	DateOfBirth string `json:"date_of_birth" gorm:"column:dob" db:"dob"`
	//CreatedAt   time.Time `json:"created_at" db:"created_at" goqu:"defaultifempty"`
	//UpdatedAt   time.Time `json:"updated_at" db:"updated_at" goqu:"defaultifempty"`
	//DeletedAt   time.Time `json:"deleted_at" db:"deleted_at" goqu:"defaultifempty"`
	Address []Address `json:"addresses" db:"-"`
}
