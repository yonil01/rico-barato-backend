package attendance

import (
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/pkg/entity/attendance"
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
