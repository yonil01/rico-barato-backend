package food

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

type ServicesFoodRepository interface {
	create(m *models.Food) error
	update(m *models.Food) error
	delete(id int) error
	getByID(id int) (*models.Food, error)
	getAll() ([]*models.Food, error)
	getFoodsByEntityId(entityId int) ([]*models.Food, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesFoodRepository {
	var s ServicesFoodRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return newFoodSqlServerRepository(db, user, txID)
	case Postgresql:
		return newFoodPsqlRepository(db, user, txID)
	case Oracle:
		return newFoodOrclRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
