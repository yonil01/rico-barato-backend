package files

import (
	"backend-comee/internal/models"
)

// Files  Model struct Files

func NewFiles(id int, entityId int, path string, typeDocument string, typeEntity int, userId string, isDelete int) *models.Files {
	return &models.Files{
		ID:           id,
		EntityId:     entityId,
		Path:         path,
		TypeDocument: typeDocument,
		TypeEntity:   typeEntity,
		UserId:       userId,
		IsDelete:     isDelete,
	}
}

func NewCreateFiles(entityId int, path string, typeDocument string, typeEntity int, userId string, isDelete int) *models.Files {
	return &models.Files{
		EntityId:     entityId,
		Path:         path,
		TypeDocument: typeDocument,
		TypeEntity:   typeEntity,
		UserId:       userId,
		IsDelete:     isDelete,
	}
}
