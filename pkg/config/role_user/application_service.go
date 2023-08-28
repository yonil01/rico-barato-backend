package role_user

import (
	"backend-ccff/internal/logger"
	"backend-ccff/internal/models"
	"fmt"
	"github.com/asaskevich/govalidator"
)

type PortsServerRoleUser interface {
	CreateRoleUser(id string, userId string, roleId string) (*RoleUser, int, error)
	UpdateRoleUser(id string, userId string, roleId string) (*RoleUser, int, error)
	DeleteRoleUser(id string) (int, error)
	GetRoleUserByID(id string) (*RoleUser, int, error)
	GetAllRoleUser() ([]*RoleUser, error)
	GetAllRoleUserByUser(id string) ([]*RoleUser, error)
}

type service struct {
	repository ServicesRoleUserRepository
	user       *models.User
	txID       string
}

func NewRoleUserService(repository ServicesRoleUserRepository, user *models.User, TxID string) PortsServerRoleUser {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateRoleUser(id string, userId string, roleId string) (*RoleUser, int, error) {
	m := NewRoleUser(id, userId, roleId)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "Dev-cff:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create RoleUser :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateRoleUser(id string, userId string, roleId string) (*RoleUser, int, error) {
	m := NewRoleUser(id, userId, roleId)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update RoleUser :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteRoleUser(id string) (int, error) {
	if !govalidator.IsUUID(id) {
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

func (s *service) GetRoleUserByID(id string) (*RoleUser, int, error) {
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

func (s *service) GetAllRoleUser() ([]*RoleUser, error) {
	return s.repository.getAll()
}

func (s *service) GetAllRoleUserByUser(id string) ([]*RoleUser, error) {
	return s.repository.getAllByUser(id)
}
