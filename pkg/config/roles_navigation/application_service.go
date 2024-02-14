package roles_navigation

import (
	"fmt"

	"backend-comee/internal/logger"
	"backend-comee/internal/models"
	"github.com/asaskevich/govalidator"
)

type PortsServerRolesNavigation interface {
	CreateRolesNavigation(id string, roleId string, navigationId string) (*RolesNavigation, int, error)
	UpdateRolesNavigation(id string, roleId string, navigationId string) (*RolesNavigation, int, error)
	DeleteRolesNavigation(id string) (int, error)
	GetRolesNavigationByID(id string) (*RolesNavigation, int, error)
	GetAllRolesNavigation() ([]*RolesNavigation, error)
}

type service struct {
	repository ServicesRolesNavigationRepository
	user       *models.User
	txID       string
}

func NewRolesNavigationService(repository ServicesRolesNavigationRepository, user *models.User, TxID string) PortsServerRolesNavigation {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateRolesNavigation(id string, roleId string, navigationId string) (*RolesNavigation, int, error) {
	m := NewRolesNavigation(id, roleId, navigationId)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create RolesNavigation :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateRolesNavigation(id string, roleId string, navigationId string) (*RolesNavigation, int, error) {
	m := NewRolesNavigation(id, roleId, navigationId)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update RolesNavigation :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteRolesNavigation(id string) (int, error) {
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

func (s *service) GetRolesNavigationByID(id string) (*RolesNavigation, int, error) {
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

func (s *service) GetAllRolesNavigation() ([]*RolesNavigation, error) {
	return s.repository.getAll()
}
