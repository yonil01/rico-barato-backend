package modules

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Model estructura de Module
type Module struct {
	ID          string    `json:"id" db:"id" valid:"required,uuid"`
	Name        string    `json:"name" db:"name" valid:"required"`
	Description string    `json:"description" db:"description" valid:"required"`
	Class       string    `json:"class" db:"class" valid:"required"`
	LicenseKey  string    `json:"license_key" db:"license_key" valid:"-"`
	Type        int       `json:"type" db:"type" valid:"required"`
	Path        string    `json:"path" db:"path" valid:"required"`
	IdUser      string    `json:"id_user" db:"id_user"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func NewModule(id string, Name string, Description string, Class string) *Module {
	return &Module{
		ID:          id,
		Name:        Name,
		Description: Description,
		Class:       Class,
	}
}

func (m *Module) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}

type ModuleRole struct {
	ID        string    `json:"id" db:"id" valid:"required,uuid"`
	RoleId    string    `json:"role_id" db:"role_id" valid:"required"`
	ElementId string    `json:"element_id" db:"element_id" valid:"required"`
	IdUser    string    `json:"id_user" db:"id_user"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
