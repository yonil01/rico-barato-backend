package events

import (
	"github.com/jmoiron/sqlx"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/logger"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/models"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesEventsRepository interface {
	create(m *Events) error
	update(m *Events) error
	delete(id string) error
	getByID(id string) (*Events, error)
	getAll() ([]*Events, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesEventsRepository {
	var s ServicesEventsRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return newEventsSqlServerRepository(db, user, txID)
	case Postgresql:
		return newEventsPsqlRepository(db, user, txID)
	case Oracle:
		return newEventsOrclRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
