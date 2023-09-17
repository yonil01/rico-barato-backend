package attendance

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Attendance  Model struct Attendance
type Attendance struct {
	ID        int       `json:"id" db:"id" valid:"-"`
	IdUser    string    `json:"id_user" db:"id_user" valid:"required"`
	IdEvent   string    `json:"id_event" db:"id_event" valid:"required"`
	IsDisable int       `json:"is_disable" db:"is_disable" valid:"-"`
	IsDelete  int       `json:"is_delete" db:"is_delete" valid:"-"`
	UserId    string    `json:"user_id" db:"user_id" valid:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewAttendance(id int, idUser string, idEvent string, isDisable int, isDelete int, userId string) *Attendance {
	return &Attendance{
		ID:        id,
		IdUser:    idUser,
		IdEvent:   idEvent,
		IsDisable: isDisable,
		IsDelete:  isDelete,
		UserId:    userId,
	}
}

func NewCreateAttendance(idUser string, idEvent string, isDisable int, isDelete int, userId string) *Attendance {
	return &Attendance{
		IdUser:    idUser,
		IdEvent:   idEvent,
		IsDisable: isDisable,
		IsDelete:  isDelete,
		UserId:    userId,
	}
}

func (m *Attendance) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
