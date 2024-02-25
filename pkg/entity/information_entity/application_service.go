package information_entity

import (
	"fmt"

	"backend-comee/internal/logger"
	"backend-comee/internal/models"
)

type PortsServerInformationEntity interface {
	CreateInformationEntity(userEntityId string, name string, description string, telephone string, mobile string, locationX string, locationY string, isBlock int, isDelete int, userId string) (*models.Entity, int, error)
	UpdateInformationEntity(id int, userEntityId string, name string, description string, telephone string, mobile string, locationX string, locationY string, isBlock int, isDelete int, userId string) (*models.Entity, int, error)
	DeleteInformationEntity(id int) (int, error)
	GetInformationEntityByID(id int) (*models.Entity, int, error)
	GetAllInformationEntity() ([]*models.Entity, error)
	GetEntityByCoordinate(long string, lat string, amount int) ([]*models.Entity, error)
}

type service struct {
	repository ServicesInformationEntityRepository
	user       *models.User
	txID       string
}

func NewInformationEntityService(repository ServicesInformationEntityRepository, user *models.User, TxID string) PortsServerInformationEntity {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateInformationEntity(userEntityId string, name string, description string, telephone string, mobile string, locationX string, locationY string, isBlock int, isDelete int, userId string) (*models.Entity, int, error) {
	m := NewCreateInformationEntity(userEntityId, name, description, telephone, mobile, locationX, locationY, isBlock, isDelete, userId)
	if valid, err := m.Valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create InformationEntity :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateInformationEntity(id int, userEntityId string, name string, description string, telephone string, mobile string, locationX string, locationY string, isBlock int, isDelete int, userId string) (*models.Entity, int, error) {
	m := NewInformationEntity(id, userEntityId, name, description, telephone, mobile, locationX, locationY, isBlock, isDelete, userId)
	if id == 0 {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return m, 15, fmt.Errorf("id is required")
	}
	if valid, err := m.Valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update InformationEntity :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteInformationEntity(id int) (int, error) {
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

func (s *service) GetInformationEntityByID(id int) (*models.Entity, int, error) {
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

func (s *service) GetAllInformationEntity() ([]*models.Entity, error) {
	return s.repository.getAll()
}

func (s *service) GetEntityByCoordinate(long string, lat string, amount int) ([]*models.Entity, error) {
	return s.repository.getEntityByCoordinate(long, lat, amount)
}
