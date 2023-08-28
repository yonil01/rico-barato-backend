package report

import (
	"backend-ccff/pkg/config/events"
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
