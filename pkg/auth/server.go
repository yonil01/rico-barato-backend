package auth

import (
	"github.com/jmoiron/sqlx"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/models"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/pkg/auth/modules"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/pkg/auth/users"
)

type ServerAuth struct {
	Users   users.PortsServerUsers
	Modules modules.PortModules
}

func NewServerAuth(db *sqlx.DB, user *models.User, txID string) *ServerAuth {
	repoDni := users.FactoryStorage(db, user, txID)
	repoModules := modules.FactoryStorage(db, user, txID)
	return &ServerAuth{
		Users:   users.NewUsersService(repoDni, user, txID),
		Modules: modules.NewModuleService(repoModules, user, txID),
	}
}
