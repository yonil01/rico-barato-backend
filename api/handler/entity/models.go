package entity

import (
	"backend-comee/internal/models"
)

type RequestEntity struct {
	UserEntityId string `json:"user_entity_id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Telephone    string `json:"telephone"`
	Mobile       string `json:"mobile"`
	LocationX    string `json:"location_x"`
	LocationY    string `json:"location_y"`
}

type ResponseAllInfoEntity struct {
	Error bool             `json:"error"`
	Data  []*models.Entity `json:"data"`
	Code  int              `json:"code"`
	Type  string           `json:"type"`
	Msg   string           `json:"msg"`
}

type ResponseInfoEntity struct {
	Error bool           `json:"error"`
	Data  *models.Entity `json:"data"`
	Code  int            `json:"code"`
	Type  string         `json:"type"`
	Msg   string         `json:"msg"`
}
