package food

import (
	"backend-comee/pkg/entity/food"
)

type RequestFood struct {
	EntityId    int    `json:"entity_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
}

type ResponseAllFood struct {
	Error bool         `json:"error"`
	Data  []*food.Food `json:"data"`
	Code  int          `json:"code"`
	Type  string       `json:"type"`
	Msg   string       `json:"msg"`
}

type ResponseFood struct {
	Error bool       `json:"error"`
	Data  *food.Food `json:"data"`
	Code  int        `json:"code"`
	Type  string     `json:"type"`
	Msg   string     `json:"msg"`
}
