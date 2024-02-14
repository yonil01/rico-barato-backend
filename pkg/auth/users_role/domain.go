package users_role

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// UsersRole  Model struct UsersRole
type UsersRole struct {
	ID        string    `json:"id" db:"id" valid:"required,uuid"`
	UserId    string    `json:"user_id" db:"user_id" valid:"required"`
	RoleId    string    `json:"role_id" db:"role_id" valid:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewUsersRole(id string, userId string, roleId string) *UsersRole {
	return &UsersRole{
		ID:     id,
		UserId: userId,
		RoleId: roleId,
	}
}

func (m *UsersRole) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
