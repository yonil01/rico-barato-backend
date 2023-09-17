package attendance

import (
	"backend-ccff/internal/logger"
	"backend-ccff/internal/models"
	"fmt"
)

type PortsServerAttendance interface {
	CreateAttendance(idUser string, idEvent string, isDisable int, isDelete int, userId string) (*Attendance, int, error)
	UpdateAttendance(id int, idUser string, idEvent string, isDisable int, isDelete int, userId string) (*Attendance, int, error)
	DeleteAttendance(id int) (int, error)
	GetAttendanceByID(id int) (*Attendance, int, error)
	GetAllAttendance() ([]*Attendance, error)
}

type service struct {
	repository ServicesAttendanceRepository
	user       *models.User
	txID       string
}

func NewAttendanceService(repository ServicesAttendanceRepository, user *models.User, TxID string) PortsServerAttendance {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateAttendance(idUser string, idEvent string, isDisable int, isDelete int, userId string) (*Attendance, int, error) {
	m := NewCreateAttendance(idUser, idEvent, isDisable, isDelete, userId)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "Dev-cff:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Attendance :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateAttendance(id int, idUser string, idEvent string, isDisable int, isDelete int, userId string) (*Attendance, int, error) {
	m := NewAttendance(id, idUser, idEvent, isDisable, isDelete, userId)
	if id == 0 {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return m, 15, fmt.Errorf("id is required")
	}
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Attendance :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteAttendance(id int) (int, error) {
	if id == 0 {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return 15, fmt.Errorf("id is required")
	}

	if err := s.repository.delete(id); err != nil {
		if err.Error() == "Dev-cff:108" {
			return 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't update row:", err)
		return 20, err
	}
	return 28, nil
}

func (s *service) GetAttendanceByID(id int) (*Attendance, int, error) {
	if id == 0 {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return nil, 15, fmt.Errorf("id is required")
	}
	m, err := s.repository.getByID(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *service) GetAllAttendance() ([]*Attendance, error) {
	return s.repository.getAll()
}
