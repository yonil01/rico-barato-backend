package auth

import (
	"backend-comee/internal/models"
	"backend-comee/pkg/auth/logger_user"
	"backend-comee/pkg/auth/modules"
	"backend-comee/pkg/auth/roles"
	"backend-comee/pkg/auth/user_entity"
	"backend-comee/pkg/auth/users"
	"github.com/jmoiron/sqlx"
)

type ServerAuth struct {
	Users      users.PortsServerUser
	Modules    modules.PortModules
	Roles      roles.PortsServerRoles
	UserEntity user_entity.PortsServerUserEntity
	LoggerUser logger_user.PortsServerLoggerUser
}

func NewServerAuth(db *sqlx.DB, user *models.User, txID string) *ServerAuth {
	repoDni := users.FactoryStorage(db, user, txID)
	repoModules := modules.FactoryStorage(db, user, txID)
	repoRoles := roles.FactoryStorage(db, user, txID)
	repoUserEntity := user_entity.FactoryStorage(db, user, txID)
	repoLoggerUser := logger_user.FactoryStorage(db, user, txID)
	return &ServerAuth{
		Users:      users.NewUserService(repoDni, user, txID),
		Modules:    modules.NewModuleService(repoModules, user, txID),
		Roles:      roles.NewRolesService(repoRoles, user, txID),
		UserEntity: user_entity.NewUserEntityService(repoUserEntity, user, txID),
		LoggerUser: logger_user.NewLoggerUserService(repoLoggerUser, user, txID),
	}
}
