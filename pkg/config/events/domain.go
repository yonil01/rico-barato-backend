package events

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Events  Model struct Events
type Events struct {
	ID          string    `json:"id" db:"id" valid:"required,uuid"`
	Name        string    `json:"name" db:"name" valid:"required"`
	Description string    `json:"description" db:"description" valid:"required"`
	EventDate   time.Time `json:"event_date" db:"event_date" valid:"-"`
	IsDisable   int       `json:"is_disable" db:"is_disable" valid:"-"`
	IsDelete    int       `json:"is_delete" db:"is_delete" valid:"-"`
	UserId      string    `json:"user_id" db:"user_id" valid:"required"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func NewEvents(id string, name string, description string, eventDate time.Time, isDisable int, isDelete int, userId string) *Events {
	return &Events{
		ID:          id,
		Name:        name,
		Description: description,
		EventDate:   eventDate,
		IsDisable:   isDisable,
		IsDelete:    isDelete,
		UserId:      userId,
	}
}

func (m *Events) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
