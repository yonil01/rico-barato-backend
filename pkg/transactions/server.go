package report

import (
	"github.com/jmoiron/sqlx"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/models"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/pkg/transactions/report"
)

type ServerReport struct {
	Event report.PortsServerReport
}

func NewServerReport(db *sqlx.DB, user *models.User, txID string) *ServerReport {
	repoReport := report.FactoryStorage(db, user, txID)
	return &ServerReport{
		Event: report.NewReportService(repoReport, user, txID),
	}
}
