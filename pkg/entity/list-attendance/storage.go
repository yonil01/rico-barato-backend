package list_ListAttendance

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

type ServicesListAttendanceRepository interface {
	create(m *ListAttendance) error
	update(m *ListAttendance) error
	delete(id int) error
	getByID(id int) (*ListAttendance, error)
	getAll() ([]*ListAttendance, error)
	getListAttendanceUser() ([]*ListAttendance, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesListAttendanceRepository {
	var s ServicesListAttendanceRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return newListAttendanceSqlServerRepository(db, user, txID)
	case Postgresql:
		return newListAttendancePsqlRepository(db, user, txID)
	case Oracle:
		return newListAttendanceOrclRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
