package events

import (
	"backend-ccff/pkg/config/events"
	"time"
)

type RequestEvent struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	EventDate   time.Time `json:"event_date"`
}

type ResponseAllEvent struct {
	Error bool             `json:"error"`
	Data  []*events.Events `json:"data"`
	Code  int              `json:"code"`
	Type  string           `json:"type"`
	Msg   string           `json:"msg"`
}

type ResponseEvent struct {
	Error bool           `json:"error"`
	Data  *events.Events `json:"data"`
	Code  int            `json:"code"`
	Type  string         `json:"type"`
	Msg   string         `json:"msg"`
}
