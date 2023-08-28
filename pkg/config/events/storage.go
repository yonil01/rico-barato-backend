package events

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
