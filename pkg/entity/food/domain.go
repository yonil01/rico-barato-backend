package food

import (
	"backend-comee/pkg/doc/files"
	"time"

	"github.com/asaskevich/govalidator"
)

// Food  Model struct Food
type Food struct {
	ID          int            `json:"id" db:"id" valid:"-"`
	EntityId    int            `json:"entity_id" db:"entity_id" valid:"required"`
	Name        string         `json:"name" db:"name" valid:"required"`
	Description string         `json:"description" db:"description" valid:"required"`
	Price       string         `json:"price" db:"price" valid:"required"`
	Status      int            `json:"status" db:"status" valid:"-"`
	IsBlock     int            `json:"is_block" db:"is_block" valid:"-"`
	IsDelete    int            `json:"is_delete" db:"is_delete" valid:"-"`
	UserId      string         `json:"user_id" db:"user_id" valid:"required"`
	File        []*files.Files `json:"files" db:"files" valid:"-"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
}

func NewFood(id int, entityId int, name string, description string, price string, status int, isBlock int, isDelete int, userId string) *Food {
	return &Food{
		ID:          id,
		EntityId:    entityId,
		Name:        name,
		Description: description,
		Price:       price,
		Status:      status,
		IsBlock:     isBlock,
		IsDelete:    isDelete,
		UserId:      userId,
	}
}

func NewCreateFood(entityId int, name string, description string, price string, status int, isBlock int, isDelete int, userId string) *Food {
	return &Food{
		EntityId:    entityId,
		Name:        name,
		Description: description,
		Price:       price,
		Status:      status,
		IsBlock:     isBlock,
		IsDelete:    isDelete,
		UserId:      userId,
	}
}

func (m *Food) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
