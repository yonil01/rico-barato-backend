package comments

import (
	"fmt"

	"backend-comee/internal/logger"
	"backend-comee/internal/models"
)

type PortsServerComment interface {
	CreateComment(userId string, entityId int, value string, start int, isDelete int) (*Comment, int, error)
	UpdateComment(id int, userId string, entityId int, value string, start int, isDelete int) (*Comment, int, error)
	DeleteComment(id int) (int, error)
	GetCommentByID(id int) (*Comment, int, error)
	GetAllComment() ([]*Comment, error)
	GetCommentByEntityID(id int) ([]*Comment, error)
}

type service struct {
	repository ServicesCommentRepository
	user       *models.User
	txID       string
}

func NewCommentService(repository ServicesCommentRepository, user *models.User, TxID string) PortsServerComment {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateComment(userId string, entityId int, value string, start int, isDelete int) (*Comment, int, error) {
	m := NewCreateComment(userId, entityId, value, start, isDelete)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Comment :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateComment(id int, userId string, entityId int, value string, start int, isDelete int) (*Comment, int, error) {
	m := NewComment(id, userId, entityId, value, start, isDelete)
	if id == 0 {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return m, 15, fmt.Errorf("id is required")
	}
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Comment :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteComment(id int) (int, error) {
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

func (s *service) GetCommentByID(id int) (*Comment, int, error) {
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

func (s *service) GetAllComment() ([]*Comment, error) {
	return s.repository.getAll()
}

func (s *service) GetCommentByEntityID(id int) ([]*Comment, error) {
	return s.repository.getCommentByEntityID(id)
}
