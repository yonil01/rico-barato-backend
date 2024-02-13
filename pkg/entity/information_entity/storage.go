package information_entity

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

type ServicesInformationEntityRepository interface {
	create(m *InformationEntity) error
	update(m *InformationEntity) error
	delete(id int) error
	getByID(id int) (*InformationEntity, error)
	getAll() ([]*InformationEntity, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesInformationEntityRepository {
	var s ServicesInformationEntityRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return newInformationEntitySqlServerRepository(db, user, txID)
	case Postgresql:
		return newInformationEntityPsqlRepository(db, user, txID)
	case Oracle:
		return newInformationEntityOrclRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
