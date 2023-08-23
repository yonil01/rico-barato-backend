package report

import "github.com/asaskevich/govalidator"

type Report struct {
	Procedure  string                 `json:"procedure"`
	Parameters map[string]interface{} `json:"parameters"`
}

func NewReport(Procedure string, Parameters map[string]interface{}) *Report {
	return &Report{
		Procedure:  Procedure,
		Parameters: Parameters,
	}
}

func (m *Report) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
