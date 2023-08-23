package Roles

import (
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/pkg/auth/roles"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/pkg/config/role_user"
)

type RequestRoles struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ResponseAllRoles struct {
	Error bool          `json:"error"`
	Data  []*roles.Role `json:"data"`
	Code  int           `json:"code"`
	Type  string        `json:"type"`
	Msg   string        `json:"msg"`
}

type ResponseRoles struct {
	Error bool        `json:"error"`
	Data  *roles.Role `json:"data"`
	Code  int         `json:"code"`
	Type  string      `json:"type"`
	Msg   string      `json:"msg"`
}

type RequestRoleUser struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
	RoleId string `json:"role_id"`
}

type ResponseAllRolesUser struct {
	Error bool                  `json:"error"`
	Data  []*role_user.RoleUser `json:"data"`
	Code  int                   `json:"code"`
	Type  string                `json:"type"`
	Msg   string                `json:"msg"`
}

type ResponseRolesUser struct {
	Error bool                `json:"error"`
	Data  *role_user.RoleUser `json:"data"`
	Code  int                 `json:"code"`
	Type  string              `json:"type"`
	Msg   string              `json:"msg"`
}
