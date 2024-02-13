package users

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

type ServicesUserRepository interface {
	create(m *User) error
	update(m *User) error
	delete(id string) error
	getByID(id string) (*User, error)
	getAll() ([]*User, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesUserRepository {
	var s ServicesUserRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return newUserSqlServerRepository(db, user, txID)
	case Postgresql:
		return newUserPsqlRepository(db, user, txID)
	case Oracle:
		return newUserOrclRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
