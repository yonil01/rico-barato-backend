package files

import (
	"backend-comee/pkg/doc/files"
)

type RequestFiles struct {
	EntityId     int    `json:"entity_id"`
	Path         string `json:"path"`
	TypeDocument string `json:"type_document"`
	TypeEntity   int    `json:"type_entity"`
}

type ResponseAllFiles struct {
	Error bool           `json:"error"`
	Data  []*files.Files `json:"data"`
	Code  int            `json:"code"`
	Type  string         `json:"type"`
	Msg   string         `json:"msg"`
}

type ResponseFiles struct {
	Error bool          `json:"error"`
	Data  *RequestFiles `json:"data"`
	Code  int           `json:"code"`
	Type  string        `json:"type"`
	Msg   string        `json:"msg"`
}
