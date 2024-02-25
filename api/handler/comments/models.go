package comments

import (
	"backend-comee/pkg/doc/comments"
)

type RequestComments struct {
	UserId   string `json:"user_id"`
	EntityId int    `json:"entity_id"`
	Value    string `json:"value"`
	Start    int    `json:"start"`
}

type ResponseAllComments struct {
	Error bool                `json:"error"`
	Data  []*comments.Comment `json:"data"`
	Code  int                 `json:"code"`
	Type  string              `json:"type"`
	Msg   string              `json:"msg"`
}

type ResponseComments struct {
	Error bool              `json:"error"`
	Data  *comments.Comment `json:"data"`
	Code  int               `json:"code"`
	Type  string            `json:"type"`
	Msg   string            `json:"msg"`
}
