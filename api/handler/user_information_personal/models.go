package user_information_personal

import (
	"backend-comee/pkg/entity/user_information_personal"
)

type RequestUserInformationPersonal struct {
	UserId string `json:"user_id"`
	Gender string `json:"gender"`
	Age    string `json:"age"`
}

type ResponseUserInformationPersonal struct {
	Error bool                                               `json:"error"`
	Data  *user_information_personal.UserInformationPersonal `json:"data"`
	Code  int                                                `json:"code"`
	Type  string                                             `json:"type"`
	Msg   string                                             `json:"msg"`
}

type ResponseUserInformationPersonalAll struct {
	Error bool                                                 `json:"error"`
	Data  []*user_information_personal.UserInformationPersonal `json:"data"`
	Code  int                                                  `json:"code"`
	Type  string                                               `json:"type"`
	Msg   string                                               `json:"msg"`
}
