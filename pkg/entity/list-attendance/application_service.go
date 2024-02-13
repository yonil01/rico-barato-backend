package list_ListAttendance

import (
	"backend-ccff/internal/logger"
	"backend-ccff/internal/models"
)

type PortsServerListAttendance interface {
	GetListAttendanceUser() ([]*ListAttendance, int, error)
}

type service struct {
	repository ServicesListAttendanceRepository
	user       *models.User
	txID       string
}

func NewListAttendanceService(repository ServicesListAttendanceRepository, user *models.User, TxID string) PortsServerListAttendance {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s service) GetListAttendanceUser() ([]*ListAttendance, int, error) {
	m, err := s.repository.getListAttendanceUser()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}
