package entity

import (
	"github.com/jmoiron/sqlx"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/models"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/pkg/entity/attendance"
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
