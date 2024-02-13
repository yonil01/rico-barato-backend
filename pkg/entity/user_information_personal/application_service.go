package user_information_personal

import (
	"fmt"

	"backend-comee/internal/logger"
	"backend-comee/internal/models"
)

type PortsServerUserInformationPersonal interface {
	CreateUserInformationPersonal(userId string, gender string, age string) (*UserInformationPersonal, int, error)
	UpdateUserInformationPersonal(id int, userId string, gender string, age string) (*UserInformationPersonal, int, error)
	DeleteUserInformationPersonal(id int) (int, error)
	GetUserInformationPersonalByID(id int) (*UserInformationPersonal, int, error)
	GetAllUserInformationPersonal() ([]*UserInformationPersonal, error)
}

type service struct {
	repository ServicesUserInformationPersonalRepository
	user       *models.User
	txID       string
}

func NewUserInformationPersonalService(repository ServicesUserInformationPersonalRepository, user *models.User, TxID string) PortsServerUserInformationPersonal {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateUserInformationPersonal(userId string, gender string, age string) (*UserInformationPersonal, int, error) {
	m := NewCreateUserInformationPersonal(userId, gender, age)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create UserInformationPersonal :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateUserInformationPersonal(id int, userId string, gender string, age string) (*UserInformationPersonal, int, error) {
	m := NewUserInformationPersonal(id, userId, gender, age)
	if id == 0 {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return m, 15, fmt.Errorf("id is required")
	}
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update UserInformationPersonal :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteUserInformationPersonal(id int) (int, error) {
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

func (s *service) GetUserInformationPersonalByID(id int) (*UserInformationPersonal, int, error) {
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

func (s *service) GetAllUserInformationPersonal() ([]*UserInformationPersonal, error) {
	return s.repository.getAll()
}
