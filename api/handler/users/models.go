package users

import "backend-ccff/pkg/auth/users"

type RequestUsers struct {
	Id             string `json:"id"`
	Username       string `json:"username"`
	CodeStudent    string `json:"code_student"`
	Dni            string `json:"dni"`
	Names          string `json:"names"`
	LastnameFather string `json:"lastname_father"`
	LastnameMother string `json:"lastname_mother"`
	Email          string `json:"email"`
	Password       string `json:"password"`
}

type ResponseUsers struct {
	Error bool         `json:"error"`
	Data  *users.Users `json:"data"`
	Code  int          `json:"code"`
	Type  string       `json:"type"`
	Msg   string       `json:"msg"`
}

type ResponseUsersAll struct {
	Error bool           `json:"error"`
	Data  []*users.Users `json:"data"`
	Code  int            `json:"code"`
	Type  string         `json:"type"`
	Msg   string         `json:"msg"`
}
