package events

import (
	"backend-ccff/internal/logger"
	"backend-ccff/internal/models"
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
)

type PortsServerEvents interface {
	CreateEvents(id string, name string, description string, eventDate time.Time, isDisable int, isDelete int, userId string) (*Events, int, error)
	UpdateEvents(id string, name string, description string, eventDate time.Time, isDisable int, isDelete int, userId string) (*Events, int, error)
	DeleteEvents(id string) (int, error)
	GetEventsByID(id string) (*Events, int, error)
	GetAllEvents() ([]*Events, error)
}

type service struct {
	repository ServicesEventsRepository
	user       *models.User
	txID       string
}

func NewEventsService(repository ServicesEventsRepository, user *models.User, TxID string) PortsServerEvents {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateEvents(id string, name string, description string, eventDate time.Time, isDisable int, isDelete int, userId string) (*Events, int, error) {
	m := NewEvents(id, name, description, eventDate, isDisable, isDelete, userId)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "Dev-cff:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Events :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateEvents(id string, name string, description string, eventDate time.Time, isDisable int, isDelete int, userId string) (*Events, int, error) {
	m := NewEvents(id, name, description, eventDate, isDisable, isDelete, userId)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Events :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteEvents(id string) (int, error) {
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

func (s *service) GetEventsByID(id string) (*Events, int, error) {
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

func (s *service) GetAllEvents() ([]*Events, error) {
	return s.repository.getAll()
}
