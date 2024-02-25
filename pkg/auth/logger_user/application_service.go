package logger_user

import (
	"fmt"

	"backend-comee/internal/logger"
	"backend-comee/internal/models"
)

type PortsServerLoggerUser interface {
	CreateLoggerUser(userId string, ip string, event string) (*LoggerUser, int, error)
	UpdateLoggerUser(id int, userId string, ip string, event string) (*LoggerUser, int, error)
	DeleteLoggerUser(id int) (int, error)
	GetLoggerUserByID(id int) (*LoggerUser, int, error)
	GetAllLoggerUser() ([]*LoggerUser, error)
}

type service struct {
	repository ServicesLoggerUserRepository
	user       *models.User
	txID       string
}

func NewLoggerUserService(repository ServicesLoggerUserRepository, user *models.User, TxID string) PortsServerLoggerUser {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateLoggerUser(userId string, ip string, event string) (*LoggerUser, int, error) {
	m := NewCreateLoggerUser(userId, ip, event)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create LoggerUser :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateLoggerUser(id int, userId string, ip string, event string) (*LoggerUser, int, error) {
	m := NewLoggerUser(id, userId, ip, event)
	if id == 0 {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return m, 15, fmt.Errorf("id is required")
	}
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update LoggerUser :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteLoggerUser(id int) (int, error) {
	if id == 0 {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return 15, fmt.Errorf("id is required")
	}

	if err := s.repository.delete(id); err != nil {
		if err.Error() == "ecatch:108" {
			return 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't update row:", err)
		return 20, err
	}
	return 28, nil
}

func (s *service) GetLoggerUserByID(id int) (*LoggerUser, int, error) {
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

func (s *service) GetAllLoggerUser() ([]*LoggerUser, error) {
	return s.repository.getAll()
}
