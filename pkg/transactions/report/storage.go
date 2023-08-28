package report

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

type ServicesReportRepository interface {
	ExecuteSP(procedure string, parameters map[string]interface{}, option int) ([]map[string]interface{}, error)
	ExecuteSPBYDocumentID(procedure string, documentID int64) (int, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesReportRepository {
	var s ServicesReportRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return NewReportSqlServerRepository(db, user, txID)
	case Postgresql:
		return NewReportPsqlRepository(db, user, txID)
	case Oracle:
		fallthrough
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
