package comments

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

type ServicesCommentRepository interface {
	create(m *Comment) error
	update(m *Comment) error
	delete(id int) error
	getByID(id int) (*Comment, error)
	getAll() ([]*Comment, error)
	getCommentByEntityID(id int) ([]*Comment, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesCommentRepository {
	var s ServicesCommentRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return newCommentSqlServerRepository(db, user, txID)
	case Postgresql:
		return newCommentPsqlRepository(db, user, txID)
	case Oracle:
		return newCommentOrclRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
