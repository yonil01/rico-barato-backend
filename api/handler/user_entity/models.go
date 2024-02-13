package user_entity

import (
	"backend-comee/pkg/auth/user_entity"
)

type RequestUserEntity struct {
	Id       string `json:"id"`
	Dni      string `json:"dni"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
}

type ResponseAllInfoBasicPerson struct {
	Error bool                      `json:"error"`
	Data  []*user_entity.UserEntity `json:"data"`
	Code  int                       `json:"code"`
	Type  string                    `json:"type"`
	Msg   string                    `json:"msg"`
}

type ResponseInfoBasicPerson struct {
	Error bool                    `json:"error"`
	Data  *user_entity.UserEntity `json:"data"`
	Code  int                     `json:"code"`
	Type  string                  `json:"type"`
	Msg   string                  `json:"msg"`
}
