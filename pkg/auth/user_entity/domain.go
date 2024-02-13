package user_entity

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// UserEntity  Model struct UserEntity
type UserEntity struct {
	ID        string    `json:"id" db:"id" valid:"required,uuid"`
	Dni       string    `json:"dni" db:"dni" valid:"required"`
	Name      string    `json:"name" db:"name" valid:"required"`
	Lastname  string    `json:"lastname" db:"lastname" valid:"required"`
	Email     string    `json:"email" db:"email" valid:"required"`
	IsBlock   int       `json:"is_block" db:"is_block" valid:"-"`
	IsDelete  int       `json:"is_delete" db:"is_delete" valid:"-"`
	UserId    string    `json:"user_id" db:"user_id" valid:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewUserEntity(id string, dni string, name string, lastname string, email string, isBlock int, isDelete int, userId string) *UserEntity {
	return &UserEntity{
		ID:       id,
		Dni:      dni,
		Name:     name,
		Lastname: lastname,
		Email:    email,
		IsBlock:  isBlock,
		IsDelete: isDelete,
		UserId:   userId,
	}
}

func (m *UserEntity) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
