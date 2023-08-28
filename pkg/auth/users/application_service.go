package users

import (
	"backend-ccff/internal/logger"
	"backend-ccff/internal/models"
	"fmt"
	"strings"

	"github.com/asaskevich/govalidator"
)

type PortsServerUsers interface {
	CreateUsers(id string, username string, codeStudent string, dni string, names string, lastnameFather string, lastnameMother string, email string, password string, isDelete int, isBlock int) (*Users, int, error)
	UpdateUsers(id string, username string, codeStudent string, dni string, names string, lastnameFather string, lastnameMother string, email string, password string, isDelete int, isBlock int) (*Users, int, error)
	DeleteUsers(id string) (int, error)
	GetUsersByID(id string) (*Users, int, error)
	GetAllUsers() ([]*Users, error)
	GetUserByCodeStudent(codeStudent string) (*Users, int, error)
}

type service struct {
	repository ServicesUsersRepository
	user       *models.User
	txID       string
}

func NewUsersService(repository ServicesUsersRepository, user *models.User, TxID string) PortsServerUsers {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateUsers(id string, username string, codeStudent string, dni string, names string, lastnameFather string, lastnameMother string, email string, password string, isDelete int, isBlock int) (*Users, int, error) {
	m := NewUsers(id, username, codeStudent, dni, names, lastnameFather, lastnameMother, email, password, isDelete, isBlock)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "Dev-cff:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Users :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateUsers(id string, username string, codeStudent string, dni string, names string, lastnameFather string, lastnameMother string, email string, password string, isDelete int, isBlock int) (*Users, int, error) {
	m := NewUsers(id, username, codeStudent, dni, names, lastnameFather, lastnameMother, email, password, isDelete, isBlock)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Users :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteUsers(id string) (int, error) {
	if !govalidator.IsUUID(strings.ToLower(id)) {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id isn't uuid"))
		return 15, fmt.Errorf("id isn't uuid")
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

func (s *service) GetUsersByID(id string) (*Users, int, error) {
	if !govalidator.IsUUID(strings.ToLower(id)) {
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

func (s *service) GetUserByCodeStudent(codeStudent string) (*Users, int, error) {
	m, err := s.repository.getByCodeStudent(codeStudent)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *service) GetAllUsers() ([]*Users, error) {
	return s.repository.getAll()
}
