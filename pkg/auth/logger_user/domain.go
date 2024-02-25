package logger_user

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// LoggerUser  Model struct LoggerUser
type LoggerUser struct {
	ID        int       `json:"id" db:"id" valid:"-"`
	UserId    string    `json:"user_id" db:"user_id" valid:"required"`
	Ip        string    `json:"ip" db:"ip" valid:"required"`
	Event     string    `json:"event" db:"event" valid:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewLoggerUser(id int, userId string, ip string, event string) *LoggerUser {
	return &LoggerUser{
		ID:     id,
		UserId: userId,
		Ip:     ip,
		Event:  event,
	}
}

func NewCreateLoggerUser(userId string, ip string, event string) *LoggerUser {
	return &LoggerUser{
		UserId: userId,
		Ip:     ip,
		Event:  event,
	}
}

func (m *LoggerUser) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
