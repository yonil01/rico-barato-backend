package modules

import "backend-ccff/pkg/auth/modules"

type RequestModulesUser struct {
	Ids  []string `json:"ids"`
	Type int      `json:"type"`
}

type ResponseModules struct {
	Error bool              `json:"error"`
	Data  []*modules.Module `json:"data"`
	Code  int               `json:"code"`
	Type  string            `json:"type"`
	Msg   string            `json:"msg"`
}

type RequestModulesRole struct {
	Id        string `json:"id"`
	RoleId    string `json:"role_id"`
	ElementId string `json:"element_id"`
}

type ResponseModulesRole struct {
	Error bool                  `json:"error"`
	Data  []*modules.ModuleRole `json:"data"`
	Code  int                   `json:"code"`
	Type  string                `json:"type"`
	Msg   string                `json:"msg"`
}

type ResponseModuleRole struct {
	Error bool                `json:"error"`
	Data  *modules.ModuleRole `json:"data"`
	Code  int                 `json:"code"`
	Type  string              `json:"type"`
	Msg   string              `json:"msg"`
}
