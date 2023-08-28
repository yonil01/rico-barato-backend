package config

import (
	"backend-ccff/internal/models"
	"backend-ccff/pkg/config/events"
	"backend-ccff/pkg/config/role_user"
	"github.com/jmoiron/sqlx"
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
