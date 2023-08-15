package reniec

import "gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/pkg/reniec/dni"

type RequestReniec struct {
	Dni string `json:"dni"`
}

type ResponseReniec struct {
	Error bool               `json:"error"`
	Data  dni.ResponseReniec `json:"data"`
	Code  int                `json:"code"`
	Type  string             `json:"type"`
	Msg   string             `json:"msg"`
}
