package models

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type Food struct {
	ID          int       `json:"id" db:"id" valid:"-"`
	EntityId    int       `json:"entity_id" db:"entity_id" valid:"required"`
	Name        string    `json:"name" db:"name" valid:"required"`
	Description string    `json:"description" db:"description" valid:"required"`
	Price       string    `json:"price" db:"price" valid:"required"`
	Status      int       `json:"status" db:"status" valid:"-"`
	IsBlock     int       `json:"is_block" db:"is_block" valid:"-"`
	IsDelete    int       `json:"is_delete" db:"is_delete" valid:"-"`
	UserId      string    `json:"user_id" db:"user_id" valid:"required"`
	File        []*Files  `json:"files" db:"files" valid:"-"`
	Entity      *Entity   `json:"entity" valid:"-"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func (m *Food) Valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
