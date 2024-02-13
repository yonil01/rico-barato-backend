package users

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// User  Model struct User
type User struct {
	ID        string    `json:"id" db:"id" valid:"required,uuid"`
	Ip        string    `json:"ip" db:"ip" valid:"required"`
	Status    int       `json:"status" db:"status" valid:"-"`
	IsBlock   int       `json:"is_block" db:"is_block" valid:"-"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewUser(id string, ip string, status int, isBlock int) *User {
	return &User{
		ID:      id,
		Ip:      ip,
		Status:  status,
		IsBlock: isBlock,
	}
}

func (m *User) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
