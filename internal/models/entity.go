package models

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// InformationEntity  Model struct InformationEntity
type Entity struct {
	ID           int       `json:"id" db:"id" valid:"-"`
	UserEntityId string    `json:"user_entity_id" db:"user_entity_id" valid:"required"`
	Name         string    `json:"name" db:"name" valid:"required"`
	Description  string    `json:"description" db:"description" valid:"required"`
	Telephone    string    `json:"telephone" db:"telephone" valid:"required"`
	Mobile       string    `json:"mobile" db:"mobile" valid:"required"`
	LocationX    string    `json:"location_x" db:"location_x" valid:"required"`
	LocationY    string    `json:"location_y" db:"location_y" valid:"required"`
	IsBlock      int       `json:"is_block" db:"is_block" valid:"-"`
	IsDelete     int       `json:"is_delete" db:"is_delete" valid:"-"`
	UserId       string    `json:"user_id" db:"user_id" valid:"required"`
	File         []*Files  `json:"files" db:"files" valid:"-"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

func (m *Entity) Valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
