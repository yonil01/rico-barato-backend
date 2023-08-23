package config

import (
	"github.com/jmoiron/sqlx"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/models"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/pkg/config/events"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/pkg/config/role_user"
)

type ServerConfig struct {
	Event    events.PortsServerEvents
	RoleUser role_user.PortsServerRoleUser
}

func NewServerConfig(db *sqlx.DB, user *models.User, txID string) *ServerConfig {
	repoEvent := events.FactoryStorage(db, user, txID)
	repoRoleUser := role_user.FactoryStorage(db, user, txID)
	return &ServerConfig{
		Event:    events.NewEventsService(repoEvent, user, txID),
		RoleUser: role_user.NewRoleUserService(repoRoleUser, user, txID),
	}
}
