package users

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

type ServicesUsersRepository interface {
	create(m *Users) error
	update(m *Users) error
	delete(id string) error
	getByID(id string) (*Users, error)
	getAll() ([]*Users, error)
	getByCodeStudent(id string) (*Users, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesUsersRepository {
	var s ServicesUsersRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return newUsersSqlServerRepository(db, user, txID)
	case Postgresql:
		return newUsersPsqlRepository(db, user, txID)
	case Oracle:
		return newUsersOrclRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
