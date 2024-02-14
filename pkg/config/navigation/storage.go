package navigation

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

type ServicesNavigationRepository interface {
	create(m *Navigation) error
	update(m *Navigation) error
	delete(id string) error
	getByID(id string) (*Navigation, error)
	getAll() ([]*Navigation, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesNavigationRepository {
	var s ServicesNavigationRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return newNavigationSqlServerRepository(db, user, txID)
	case Postgresql:
		return newNavigationPsqlRepository(db, user, txID)
	case Oracle:
		return newNavigationOrclRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
