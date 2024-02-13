package list_ListAttendance

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// ListAttendance  Model struct ListAttendance
type ListAttendance struct {
	ID        int       `json:"id" db:"id" valid:"-"`
	EventName string    `json:"event_name" db:"event_name" valid:"required"`
	IsDelete  int       `json:"is_delete" db:"is_delete" valid:"-"`
	UserName  string    `json:"user_name" db:"user_name" valid:"-"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewListAttendance(id int, eventName string, isDelete int, userName string) *ListAttendance {
	return &ListAttendance{
		ID:        id,
		EventName: eventName,
		IsDelete:  isDelete,
		UserName:  userName,
	}
}

func NewCreateListAttendance(eventName string, isDelete int, userName string) *ListAttendance {
	return &ListAttendance{
		EventName: eventName,
		IsDelete:  isDelete,
		UserName:  userName,
	}
}

func (m *ListAttendance) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
