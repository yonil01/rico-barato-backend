package dni

import (
	"backend-ccff/internal/logger"
	"backend-ccff/internal/models"
	"github.com/jmoiron/sqlx"
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
