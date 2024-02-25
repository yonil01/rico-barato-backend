package information_entity

import (
	models "backend-comee/internal/models"
)

func NewInformationEntity(id int, userEntityId string, name string, description string, telephone string, mobile string, locationX string, locationY string, isBlock int, isDelete int, userId string) *models.Entity {
	return &models.Entity{
		ID:           id,
		UserEntityId: userEntityId,
		Name:         name,
		Description:  description,
		Telephone:    telephone,
		Mobile:       mobile,
		LocationX:    locationX,
		LocationY:    locationY,
		IsBlock:      isBlock,
		IsDelete:     isDelete,
		UserId:       userId,
	}
}

func NewCreateInformationEntity(userEntityId string, name string, description string, telephone string, mobile string, locationX string, locationY string, isBlock int, isDelete int, userId string) *models.Entity {
	return &models.Entity{
		UserEntityId: userEntityId,
		Name:         name,
		Description:  description,
		Telephone:    telephone,
		Mobile:       mobile,
		LocationX:    locationX,
		LocationY:    locationY,
		IsBlock:      isBlock,
		IsDelete:     isDelete,
		UserId:       userId,
	}
}
