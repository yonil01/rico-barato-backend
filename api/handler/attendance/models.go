package attendance

import (
	"backend-ccff/pkg/entity/attendance"
	list_ListAttendance "backend-ccff/pkg/entity/list-attendance"
	"time"
)

type RequestAttendance struct {
	CodeStudent    string    `json:"code_student"`
	DateAttendance time.Time `json:"date_attendance"`
	IdEvent        string    `json:"id_event"`
}

type ResponseAllAttendance struct {
	Error bool                     `json:"error"`
	Data  []*attendance.Attendance `json:"data"`
	Code  int                      `json:"code"`
	Type  string                   `json:"type"`
	Msg   string                   `json:"msg"`
}

type ResponseAttendance struct {
	Error bool                   `json:"error"`
	Data  *attendance.Attendance `json:"data"`
	Code  int                    `json:"code"`
	Type  string                 `json:"type"`
	Msg   string                 `json:"msg"`
}

type ResponseListAttendance struct {
	Error bool                                  `json:"error"`
	Data  []*list_ListAttendance.ListAttendance `json:"data"`
	Code  int                                   `json:"code"`
	Type  string                                `json:"type"`
	Msg   string                                `json:"msg"`
}
