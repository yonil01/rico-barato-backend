package navigation

import (
	"fmt"

	"backend-comee/internal/logger"
	"backend-comee/internal/models"
	"github.com/asaskevich/govalidator"
)

type PortsServerNavigation interface {
	CreateNavigation(id string, name string, description string, userId string) (*Navigation, int, error)
	UpdateNavigation(id string, name string, description string, userId string) (*Navigation, int, error)
	DeleteNavigation(id string) (int, error)
	GetNavigationByID(id string) (*Navigation, int, error)
	GetAllNavigation() ([]*Navigation, error)
}

type service struct {
	repository ServicesNavigationRepository
	user       *models.User
	txID       string
}

func NewNavigationService(repository ServicesNavigationRepository, user *models.User, TxID string) PortsServerNavigation {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateNavigation(id string, name string, description string, userId string) (*Navigation, int, error) {
	m := NewNavigation(id, name, description, userId)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Navigation :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateNavigation(id string, name string, description string, userId string) (*Navigation, int, error) {
	m := NewNavigation(id, name, description, userId)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Navigation :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteNavigation(id string) (int, error) {
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

func (s *service) GetNavigationByID(id string) (*Navigation, int, error) {
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

func (s *service) GetAllNavigation() ([]*Navigation, error) {
	return s.repository.getAll()
}
