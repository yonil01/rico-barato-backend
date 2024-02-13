package entity

import (
	"backend-comee/pkg/entity/information_entity"
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
	Error bool                                    `json:"error"`
	Data  []*information_entity.InformationEntity `json:"data"`
	Code  int                                     `json:"code"`
	Type  string                                  `json:"type"`
	Msg   string                                  `json:"msg"`
}

type ResponseInfoEntity struct {
	Error bool                                  `json:"error"`
	Data  *information_entity.InformationEntity `json:"data"`
	Code  int                                   `json:"code"`
	Type  string                                `json:"type"`
	Msg   string                                `json:"msg"`
}
