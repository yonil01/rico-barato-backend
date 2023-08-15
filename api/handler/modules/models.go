package modules

import "gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/pkg/auth/modules"

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
