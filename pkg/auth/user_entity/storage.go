package user_entity

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

type ServicesUserEntityRepository interface {
	create(m *UserEntity) error
	update(m *UserEntity) error
	delete(id string) error
	getByID(id string) (*UserEntity, error)
	getAll() ([]*UserEntity, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesUserEntityRepository {
	var s ServicesUserEntityRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return newUserEntitySqlServerRepository(db, user, txID)
	case Postgresql:
		return newUserEntityPsqlRepository(db, user, txID)
	case Oracle:
		return newUserEntityOrclRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
