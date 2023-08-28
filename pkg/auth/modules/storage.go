package modules

import (
	"backend-ccff/internal/logger"
	"backend-ccff/internal/models"
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
)

type ServicesModuleRepository interface {
	Create(m *Module) error
	Update(m *Module) error
	Delete(id string) error
	GetByID(id string) (*Module, error)
	GetAll() ([]*Module, error)
	GetModulesByRoles(roleIDs []string, ids []string, typeArg int) ([]*Module, error)
	GetModulesRole(id string) ([]*ModuleRole, error)
	DeleteModuleUser(id string) error
	CreateModuleRole(role *ModuleRole) error
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesModuleRepository {
	var s ServicesModuleRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return NewModuleSqlServerRepository(db, user, txID)
	case Postgresql:
		return NewModulePsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
