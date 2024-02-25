package files

import (
	"fmt"

	"backend-comee/internal/logger"
	"backend-comee/internal/models"
)

type PortsServerFiles interface {
	CreateFiles(entityId int, path string, typeDocument string, typeEntity int, userId string, isDelete int) (*models.Files, int, error)
	UpdateFiles(id int, entityId int, path string, typeDocument string, typeEntity int, userId string, isDelete int) (*models.Files, int, error)
	DeleteFiles(id int) (int, error)
	GetFilesByID(id int) (*models.Files, int, error)
	GetAllFiles() ([]*models.Files, error)
	GetFilesByEntityId(entityId int, typeEntity int) ([]*models.Files, error)
}

type service struct {
	repository ServicesFilesRepository
	user       *models.User
	txID       string
}

func NewFilesService(repository ServicesFilesRepository, user *models.User, TxID string) PortsServerFiles {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateFiles(entityId int, path string, typeDocument string, typeEntity int, userId string, isDelete int) (*models.Files, int, error) {
	m := NewCreateFiles(entityId, path, typeDocument, typeEntity, userId, isDelete)
	if valid, err := m.Valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Files :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateFiles(id int, entityId int, path string, typeDocument string, typeEntity int, userId string, isDelete int) (*models.Files, int, error) {
	m := NewFiles(id, entityId, path, typeDocument, typeEntity, userId, isDelete)
	if id == 0 {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return m, 15, fmt.Errorf("id is required")
	}
	if valid, err := m.Valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Files :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteFiles(id int) (int, error) {
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

func (s *service) GetFilesByID(id int) (*models.Files, int, error) {
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

func (s *service) GetAllFiles() ([]*models.Files, error) {
	return s.repository.getAll()
}

func (s *service) GetFilesByEntityId(entityId int, typeEntity int) ([]*models.Files, error) {
	return s.repository.getFilesByEntityId(entityId, typeEntity)
}
