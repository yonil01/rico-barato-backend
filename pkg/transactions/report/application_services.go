package report

import (
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/logger"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/models"
)

type PortsServerReport interface {
	ExecuteReport(procedure string, parameters map[string]interface{}, option int) ([]map[string]interface{}, error)
	ExecuteSPBYDocumentID(procedure string, dID int64) (int, error)
}

type service struct {
	repository ServicesReportRepository
	user       *models.User
	txID       string
}

func NewReportService(repository ServicesReportRepository, user *models.User, TxID string) PortsServerReport {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) ExecuteReport(procedure string, parameters map[string]interface{}, option int) ([]map[string]interface{}, error) {
	rs, err := s.repository.ExecuteSP(procedure, parameters, option)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't execute Report :", err)
		return nil, err
	}
	return rs, nil
}

func (s *service) ExecuteSPBYDocumentID(procedure string, dID int64) (int, error) {
	return s.repository.ExecuteSPBYDocumentID(procedure, dID)
}
