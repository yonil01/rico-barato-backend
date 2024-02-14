package roles_navigation

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

type ServicesRolesNavigationRepository interface {
	create(m *RolesNavigation) error
	update(m *RolesNavigation) error
	delete(id string) error
	getByID(id string) (*RolesNavigation, error)
	getAll() ([]*RolesNavigation, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesRolesNavigationRepository {
	var s ServicesRolesNavigationRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return newRolesNavigationSqlServerRepository(db, user, txID)
	case Postgresql:
		return newRolesNavigationPsqlRepository(db, user, txID)
	case Oracle:
		return newRolesNavigationOrclRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
