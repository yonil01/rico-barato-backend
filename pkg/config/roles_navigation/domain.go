package roles_navigation

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// RolesNavigation  Model struct RolesNavigation
type RolesNavigation struct {
	ID           string    `json:"id" db:"id" valid:"required,uuid"`
	RoleId       string    `json:"role_id" db:"role_id" valid:"required"`
	NavigationId string    `json:"navigation_id" db:"navigation_id" valid:"required"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

func NewRolesNavigation(id string, roleId string, navigationId string) *RolesNavigation {
	return &RolesNavigation{
		ID:           id,
		RoleId:       roleId,
		NavigationId: navigationId,
	}
}

func (m *RolesNavigation) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
