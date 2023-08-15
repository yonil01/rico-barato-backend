package dni

import (
	"github.com/jmoiron/sqlx"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/logger"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/models"
)

const (
	Postgresql = "postgres"
)

type ServicesReniecRepository interface {
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesReniecRepository {
	var s ServicesReniecRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return newReniecPsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
