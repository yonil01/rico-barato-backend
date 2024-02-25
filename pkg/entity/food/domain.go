package food

import (
	"backend-comee/internal/models"
)

// Food  Model struct Food

func NewFood(id int, entityId int, name string, description string, price string, status int, isBlock int, isDelete int, userId string) *models.Food {
	return &models.Food{
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

func NewCreateFood(entityId int, name string, description string, price string, status int, isBlock int, isDelete int, userId string) *models.Food {
	return &models.Food{
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
