package role_user

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// RoleUser  Model struct RoleUser
type RoleUser struct {
	ID        string    `json:"id" db:"id" valid:"required,uuid"`
	UserId    string    `json:"user_id" db:"user_id" valid:"required"`
	RoleId    string    `json:"role_id" db:"role_id" valid:"required"`
	IdUser    string    `json:"id_user" db:"id_user" valid:"-"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewRoleUser(id string, userId string, roleId string) *RoleUser {
	return &RoleUser{
		ID:     id,
		UserId: userId,
		RoleId: roleId,
	}
}

func (m *RoleUser) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
