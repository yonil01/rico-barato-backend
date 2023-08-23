package report

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
