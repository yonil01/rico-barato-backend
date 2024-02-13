package users

import (
	"fmt"

	"backend-comee/internal/logger"
	"backend-comee/internal/models"
	"github.com/asaskevich/govalidator"
)

type PortsServerUser interface {
	CreateUser(id string, ip string, status int, isBlock int) (*User, int, error)
	UpdateUser(id string, ip string, status int, isBlock int) (*User, int, error)
	DeleteUser(id string) (int, error)
	GetUserByID(id string) (*User, int, error)
	GetAllUser() ([]*User, error)
}

type service struct {
	repository ServicesUserRepository
	user       *models.User
	txID       string
}

func NewUserService(repository ServicesUserRepository, user *models.User, TxID string) PortsServerUser {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateUser(id string, ip string, status int, isBlock int) (*User, int, error) {
	m := NewUser(id, ip, status, isBlock)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create User :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateUser(id string, ip string, status int, isBlock int) (*User, int, error) {
	m := NewUser(id, ip, status, isBlock)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update User :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteUser(id string) (int, error) {
	if !govalidator.IsUUID(id) {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id isn't uuid"))
		return 15, fmt.Errorf("id isn't uuid")
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

func (s *service) GetUserByID(id string) (*User, int, error) {
	if !govalidator.IsUUID(id) {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id isn't uuid"))
		return nil, 15, fmt.Errorf("id isn't uuid")
	}
	m, err := s.repository.getByID(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *service) GetAllUser() ([]*User, error) {
	return s.repository.getAll()
}
