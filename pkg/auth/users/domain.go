package users

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Users  Model struct Users
type Users struct {
	ID             string    `json:"id" db:"id" valid:"required,uuid"`
	Username       string    `json:"username" db:"username" valid:"required"`
	CodeStudent    string    `json:"code_student" db:"code_student" valid:"required"`
	Dni            string    `json:"dni" db:"dni" valid:"required"`
	Names          string    `json:"names" db:"names" valid:"required"`
	LastnameFather string    `json:"lastname_father" db:"lastname_father" valid:"required"`
	LastnameMother string    `json:"lastname_mother" db:"lastname_mother" valid:"required"`
	Email          string    `json:"email" db:"email" valid:"required"`
	Password       string    `json:"password" db:"password" valid:"required"`
	IsDelete       int       `json:"is_delete" db:"is_delete" valid:"-"`
	IsBlock        int       `json:"is_block" db:"is_block" valid:"-"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

func NewUsers(id string, username string, codeStudent string, dni string, names string, lastnameFather string, lastnameMother string, email string, password string, isDelete int, isBlock int) *Users {
	return &Users{
		ID:             id,
		Username:       username,
		CodeStudent:    codeStudent,
		Dni:            dni,
		Names:          names,
		LastnameFather: lastnameFather,
		LastnameMother: lastnameMother,
		Email:          email,
		Password:       password,
		IsDelete:       isDelete,
		IsBlock:        isBlock,
	}
}

func (m *Users) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
