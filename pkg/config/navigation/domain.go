package navigation

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Navigation  Model struct Navigation
type Navigation struct {
	ID          string    `json:"id" db:"id" valid:"required,uuid"`
	Name        string    `json:"name" db:"name" valid:"required"`
	Description string    `json:"description" db:"description" valid:"required"`
	UserId      string    `json:"user_id" db:"user_id" valid:"required"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func NewNavigation(id string, name string, description string, userId string) *Navigation {
	return &Navigation{
		ID:          id,
		Name:        name,
		Description: description,
		UserId:      userId,
	}
}

func (m *Navigation) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
