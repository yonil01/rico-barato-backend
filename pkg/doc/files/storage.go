package files

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

type ServicesFilesRepository interface {
	create(m *Files) error
	update(m *Files) error
	delete(id int) error
	getByID(id int) (*Files, error)
	getAll() ([]*Files, error)
	getFilesByEntityId(entityId int, typeEntity int) ([]*Files, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesFilesRepository {
	var s ServicesFilesRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return newFilesSqlServerRepository(db, user, txID)
	case Postgresql:
		return newFilesPsqlRepository(db, user, txID)
	case Oracle:
		return newFilesOrclRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}