package files

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Files  Model struct Files
type Files struct {
	ID           int       `json:"id" db:"id" valid:"-"`
	EntityId     int       `json:"entity_id" db:"entity_id" valid:"required"`
	Path         string    `json:"path" db:"path" valid:"required"`
	TypeDocument string    `json:"type_document" db:"type_document" valid:"required"`
	TypeEntity   int       `json:"type_entity" db:"type_entity" valid:"required"`
	UserId       string    `json:"user_id" db:"user_id" valid:"required"`
	IsDelete     int       `json:"is_delete" db:"is_delete" valid:"-"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

func NewFiles(id int, entityId int, path string, typeDocument string, typeEntity int, userId string, isDelete int) *Files {
	return &Files{
		ID:           id,
		EntityId:     entityId,
		Path:         path,
		TypeDocument: typeDocument,
		TypeEntity:   typeEntity,
		UserId:       userId,
		IsDelete:     isDelete,
	}
}

func NewCreateFiles(entityId int, path string, typeDocument string, typeEntity int, userId string, isDelete int) *Files {
	return &Files{
		EntityId:     entityId,
		Path:         path,
		TypeDocument: typeDocument,
		TypeEntity:   typeEntity,
		UserId:       userId,
		IsDelete:     isDelete,
	}
}

func (m *Files) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
