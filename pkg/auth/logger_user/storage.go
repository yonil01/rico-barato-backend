package logger_user

import (
	"github.com/jmoiron/sqlx"

	"backend-comee/internal/logger"
	"backend-comee/internal/models"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesLoggerUserRepository interface {
	create(m *LoggerUser) error
	update(m *LoggerUser) error
	delete(id int) error
	getByID(id int) (*LoggerUser, error)
	getAll() ([]*LoggerUser, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesLoggerUserRepository {
	var s ServicesLoggerUserRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return newLoggerUserSqlServerRepository(db, user, txID)
	case Postgresql:
		return newLoggerUserPsqlRepository(db, user, txID)
	case Oracle:
		return newLoggerUserOrclRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
