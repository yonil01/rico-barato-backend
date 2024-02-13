package entity

import (
	"backend-ccff/internal/models"
	"backend-ccff/pkg/entity/attendance"
	list_ListAttendance "backend-ccff/pkg/entity/list-attendance"
	"github.com/jmoiron/sqlx"
)

type ServerEntity struct {
	Attendance     attendance.PortsServerAttendance
	ListAttendance list_ListAttendance.PortsServerListAttendance
}

func NewServerEntity(db *sqlx.DB, user *models.User, txID string) *ServerEntity {
	repoAttendance := attendance.FactoryStorage(db, user, txID)
	repoListAttendance := list_ListAttendance.FactoryStorage(db, user, txID)
	return &ServerEntity{
		Attendance:     attendance.NewAttendanceService(repoAttendance, user, txID),
		ListAttendance: list_ListAttendance.NewListAttendanceService(repoListAttendance, user, txID),
	}
}
