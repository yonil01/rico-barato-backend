package attendance

import (
	"backend-ccff/internal/logger"
	"backend-ccff/internal/models"
	"github.com/jmoiron/sqlx"
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
