package report

import (
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/pkg/config/events"
)

type RequestReport struct {
	Procedure  string      `json:"procedure"`
	Parameters interface{} `json:"parameters"`
}

type ResponseAllEvent struct {
	Error bool             `json:"error"`
	Data  []*events.Events `json:"data"`
	Code  int              `json:"code"`
	Type  string           `json:"type"`
	Msg   string           `json:"msg"`
}

type ResponseReport struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
	Code  int         `json:"code"`
	Type  string      `json:"type"`
	Msg   string      `json:"msg"`
}
