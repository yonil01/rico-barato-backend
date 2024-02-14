package users_role

import (
	"fmt"

	"backend-comee/internal/logger"
	"backend-comee/internal/models"
	"github.com/asaskevich/govalidator"
)

type PortsServerUsersRole interface {
	CreateUsersRole(id string, userId string, roleId string) (*UsersRole, int, error)
	UpdateUsersRole(id string, userId string, roleId string) (*UsersRole, int, error)
	DeleteUsersRole(id string) (int, error)
	GetUsersRoleByID(id string) (*UsersRole, int, error)
	GetAllUsersRole() ([]*UsersRole, error)
}

type service struct {
	repository ServicesUsersRoleRepository
	user       *models.User
	txID       string
}

func NewUsersRoleService(repository ServicesUsersRoleRepository, user *models.User, TxID string) PortsServerUsersRole {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateUsersRole(id string, userId string, roleId string) (*UsersRole, int, error) {
	m := NewUsersRole(id, userId, roleId)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create UsersRole :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateUsersRole(id string, userId string, roleId string) (*UsersRole, int, error) {
	m := NewUsersRole(id, userId, roleId)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update UsersRole :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteUsersRole(id string) (int, error) {
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

func (s *service) GetUsersRoleByID(id string) (*UsersRole, int, error) {
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

func (s *service) GetAllUsersRole() ([]*UsersRole, error) {
	return s.repository.getAll()
}
