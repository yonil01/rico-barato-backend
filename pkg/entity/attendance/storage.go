package attendance

import (
	"github.com/jmoiron/sqlx"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/logger"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/models"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesAttendanceRepository interface {
	create(m *Attendance) error
	update(m *Attendance) error
	delete(id int) error
	getByID(id int) (*Attendance, error)
	getAll() ([]*Attendance, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesAttendanceRepository {
	var s ServicesAttendanceRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return newAttendanceSqlServerRepository(db, user, txID)
	case Postgresql:
		return newAttendancePsqlRepository(db, user, txID)
	case Oracle:
		return newAttendanceOrclRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
