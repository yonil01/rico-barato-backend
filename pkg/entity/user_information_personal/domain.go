package user_information_personal

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// UserInformationPersonal  Model struct UserInformationPersonal
type UserInformationPersonal struct {
	ID        int       `json:"id" db:"id" valid:"-"`
	UserId    string    `json:"user_id" db:"user_id" valid:"required"`
	Gender    string    `json:"gender" db:"gender" valid:"required"`
	Age       string    `json:"age" db:"age" valid:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewUserInformationPersonal(id int, userId string, gender string, age string) *UserInformationPersonal {
	return &UserInformationPersonal{
		ID:     id,
		UserId: userId,
		Gender: gender,
		Age:    age,
	}
}

func NewCreateUserInformationPersonal(userId string, gender string, age string) *UserInformationPersonal {
	return &UserInformationPersonal{
		UserId: userId,
		Gender: gender,
		Age:    age,
	}
}

func (m *UserInformationPersonal) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
