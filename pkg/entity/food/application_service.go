package food

import (
	"fmt"

	"backend-comee/internal/logger"
	"backend-comee/internal/models"
)

type PortsServerFood interface {
	CreateFood(entityId int, name string, description string, price string, status int, isBlock int, isDelete int, userId string) (*models.Food, int, error)
	UpdateFood(id int, entityId int, name string, description string, price string, status int, isBlock int, isDelete int, userId string) (*models.Food, int, error)
	DeleteFood(id int) (int, error)
	GetFoodByID(id int) (*models.Food, int, error)
	GetAllFood() ([]*models.Food, error)
	GetFoodsByEntityId(entityId int) ([]*models.Food, error)
}

type service struct {
	repository ServicesFoodRepository
	user       *models.User
	txID       string
}

func NewFoodService(repository ServicesFoodRepository, user *models.User, TxID string) PortsServerFood {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateFood(entityId int, name string, description string, price string, status int, isBlock int, isDelete int, userId string) (*models.Food, int, error) {
	m := NewCreateFood(entityId, name, description, price, status, isBlock, isDelete, userId)
	if valid, err := m.Valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Food :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateFood(id int, entityId int, name string, description string, price string, status int, isBlock int, isDelete int, userId string) (*models.Food, int, error) {
	m := NewFood(id, entityId, name, description, price, status, isBlock, isDelete, userId)
	if id == 0 {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return m, 15, fmt.Errorf("id is required")
	}
	if valid, err := m.Valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Food :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteFood(id int) (int, error) {
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

func (s *service) GetFoodByID(id int) (*models.Food, int, error) {
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

func (s *service) GetAllFood() ([]*models.Food, error) {
	return s.repository.getAll()
}

func (s *service) GetFoodsByEntityId(entityId int) ([]*models.Food, error) {
	return s.repository.getFoodsByEntityId(entityId)
}
