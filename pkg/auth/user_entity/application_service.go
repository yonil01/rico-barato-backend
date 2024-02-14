package user_entity

import (
	"fmt"

	"backend-comee/internal/logger"
	"backend-comee/internal/models"
	"github.com/asaskevich/govalidator"
)

type PortsServerUserEntity interface {
	CreateUserEntity(id string, dni string, name string, lastname string, email string, password string, isBlock int, isDelete int, userId string) (*UserEntity, int, error)
	UpdateUserEntity(id string, dni string, name string, lastname string, email string, password string, isBlock int, isDelete int, userId string) (*UserEntity, int, error)
	DeleteUserEntity(id string) (int, error)
	GetUserEntityByID(id string) (*UserEntity, int, error)
	GetAllUserEntity() ([]*UserEntity, error)
}

type service struct {
	repository ServicesUserEntityRepository
	user       *models.User
	txID       string
}

func NewUserEntityService(repository ServicesUserEntityRepository, user *models.User, TxID string) PortsServerUserEntity {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateUserEntity(id string, dni string, name string, lastname string, email string, password string, isBlock int, isDelete int, userId string) (*UserEntity, int, error) {
	m := NewUserEntity(id, dni, name, lastname, email, password, isBlock, isDelete, userId)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create UserEntity :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateUserEntity(id string, dni string, name string, lastname string, email string, password string, isBlock int, isDelete int, userId string) (*UserEntity, int, error) {
	m := NewUserEntity(id, dni, name, lastname, email, password, isBlock, isDelete, userId)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update UserEntity :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteUserEntity(id string) (int, error) {
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

func (s *service) GetUserEntityByID(id string) (*UserEntity, int, error) {
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

func (s *service) GetAllUserEntity() ([]*UserEntity, error) {
	return s.repository.getAll()
}
