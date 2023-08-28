package modules

import (
	"backend-ccff/internal/logger"
	"backend-ccff/internal/models"
	"fmt"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Service struct {
	repository ServicesModuleRepository
	user       *models.User
	txID       string
}

type PortModules interface {
	GetModulesByRoles(roleIDs []string, ids []string, typeArg int) ([]*Module, error)
	GetAllModule() ([]*Module, error)
	GetModulesRole(roleId string) ([]*ModuleRole, error)
	DeleteModulesRole(id string) (int, error)
	CreateModulesRole(id string, roleId string, elementId string) (*ModuleRole, int, error)
}

func NewModuleService(repository ServicesModuleRepository, user *models.User, TxID string) Service {
	return Service{repository: repository, user: user, txID: TxID}
}

func (s Service) CreateModule(id string, Name string, Description string, Class string) (*Module, int, error) {
	m := NewModule(id, Name, Description, Class)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.Create(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't create Module :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s Service) UpdateModule(id string, Name string, Description string, Class string) (*Module, int, error) {
	m := NewModule(id, Name, Description, Class)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.Update(m); err != nil {
		if err.Error() == "Dev-cff:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't update Module :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s Service) DeleteModule(id string) (int, error) {
	if !govalidator.IsUUID(id) {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id isn't uuid"))
		return 15, fmt.Errorf("id isn't uuid")
	}

	if err := s.repository.Delete(id); err != nil {
		if err.Error() == "Dev-cff:108" {
			return 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't update row:", err)
		return 20, err
	}
	return 28, nil
}

func (s Service) GetModuleByID(id string) (*Module, int, error) {
	if !govalidator.IsUUID(id) {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id isn't uuid"))
		return nil, 15, fmt.Errorf("id isn't uuid")
	}
	m, err := s.repository.GetByID(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't getByID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s Service) GetAllModule() ([]*Module, error) {
	return s.repository.GetAll()
}

func (s Service) GetModulesByRoles(roleIDs []string, ids []string, typeArg int) ([]*Module, error) {
	return s.repository.GetModulesByRoles(roleIDs, ids, typeArg)
}

func (s Service) GetModulesRole(roleId string) ([]*ModuleRole, error) {
	return s.repository.GetModulesRole(roleId)
}

func (s Service) DeleteModulesRole(id string) (int, error) {
	if !govalidator.IsUUID(strings.ToLower(id)) {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id isn't uuid"))
		return 15, fmt.Errorf("id isn't uuid")
	}

	if err := s.repository.DeleteModuleUser(id); err != nil {
		if err.Error() == "Dev-cff:108" {
			return 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't update row:", err)
		return 20, err
	}
	return 28, nil
}

func (s Service) CreateModulesRole(id string, roleId string, elementId string) (*ModuleRole, int, error) {
	if !govalidator.IsUUID(strings.ToLower(id)) {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id isn't uuid"))
		return nil, 15, fmt.Errorf("id isn't uuid")
	}

	newModuleRole := ModuleRole{
		ID:        id,
		RoleId:    roleId,
		ElementId: elementId,
	}

	if err := s.repository.CreateModuleRole(&newModuleRole); err != nil {
		if err.Error() == "Dev-cff:108" {
			return nil, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't update row:", err)
		return nil, 20, err
	}
	return &newModuleRole, 28, nil
}
