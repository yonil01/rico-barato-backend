package role_user

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

type ServicesRoleUserRepository interface {
	create(m *RoleUser) error
	update(m *RoleUser) error
	delete(id string) error
	getByID(id string) (*RoleUser, error)
	getAll() ([]*RoleUser, error)
	getAllByUser(id string) ([]*RoleUser, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesRoleUserRepository {
	var s ServicesRoleUserRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return newRoleUserSqlServerRepository(db, user, txID)
	case Postgresql:
		return newRoleUserPsqlRepository(db, user, txID)
	case Oracle:
		return newRoleUserOrclRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
