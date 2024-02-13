package users

import (
	"backend-comee/pkg/auth/users"
)

type RequestUsers struct {
	Id string `json:"id"`
	Ip string `json:"ip"`
}

type ResponseUsers struct {
	Error bool        `json:"error"`
	Data  *users.User `json:"data"`
	Code  int         `json:"code"`
	Type  string      `json:"type"`
	Msg   string      `json:"msg"`
}

type ResponseUsersAll struct {
	Error bool          `json:"error"`
	Data  []*users.User `json:"data"`
	Code  int           `json:"code"`
	Type  string        `json:"type"`
	Msg   string        `json:"msg"`
}
