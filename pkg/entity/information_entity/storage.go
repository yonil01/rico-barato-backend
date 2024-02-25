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
	create(m *models.Entity) error
	update(m *models.Entity) error
	delete(id int) error
	getByID(id int) (*models.Entity, error)
	getAll() ([]*models.Entity, error)
	getEntityByCoordinate(locationX string, locationY string, amount int) ([]*models.Entity, error)
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
