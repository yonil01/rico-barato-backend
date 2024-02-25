package models

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type Files struct {
	ID           int       `json:"id" db:"id" valid:"-"`
	EntityId     int       `json:"entity_id" db:"entity_id" valid:"required"`
	Path         string    `json:"path" db:"path" valid:"required"`
	TypeDocument string    `json:"type_document" db:"type_document" valid:"required"`
	TypeEntity   int       `json:"type_entity" db:"type_entity" valid:"required"`
	UserId       string    `json:"user_id" db:"user_id" valid:"required"`
	IsDelete     int       `json:"is_delete" db:"is_delete" valid:"-"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

func (m *Files) Valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
