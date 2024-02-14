package users_role

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

type ServicesUsersRoleRepository interface {
	create(m *UsersRole) error
	update(m *UsersRole) error
	delete(id string) error
	getByID(id string) (*UsersRole, error)
	getAll() ([]*UsersRole, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesUsersRoleRepository {
	var s ServicesUsersRoleRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return newUsersRoleSqlServerRepository(db, user, txID)
	case Postgresql:
		return newUsersRolePsqlRepository(db, user, txID)
	case Oracle:
		return newUsersRoleOrclRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
