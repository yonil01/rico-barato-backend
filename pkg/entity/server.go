package entity

import (
	"backend-ccff/internal/models"
	"backend-ccff/pkg/entity/attendance"
	"github.com/jmoiron/sqlx"
)

type ServerEntity struct {
	Attendance attendance.PortsServerAttendance
}

func NewServerEntity(db *sqlx.DB, user *models.User, txID string) *ServerEntity {
	repoAttendance := attendance.FactoryStorage(db, user, txID)
	return &ServerEntity{
		Attendance: attendance.NewAttendanceService(repoAttendance, user, txID),
	}
}
