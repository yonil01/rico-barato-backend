package dni

import (
	"encoding/json"
	"fmt"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/env"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/logger"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/models"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/ws"
)

type PortsServerReniec interface {
	GetConsultReniecByDni(dni string) (ResponseReniec, int, error)
}

type service struct {
	repository ServicesReniecRepository
	user       *models.User
	txID       string
}

func NewReniecService(repository ServicesReniecRepository, user *models.User, txID string) PortsServerReniec {
	return &service{
		repository: repository,
		user:       user,
		txID:       txID,
	}
}

func (s *service) GetConsultReniecByDni(dni string) (ResponseReniec, int, error) {
	c := env.NewConfiguration()

	urlReniec := fmt.Sprintf("%s%s%s", c.External.Reniec, dni, c.External.Credential)

	var resp ResponseReniec
	response, cod, err := ws.ConsumeWS(nil, urlReniec, "GET")
	if err != nil {
		logger.Error.Printf("Couldn't insert ConsumeWS: %v", err)
		return resp, cod, err
	}

	if cod != 200 {
		logger.Error.Println("respuesta servicio no exitosa: ", cod)
		return resp, cod, err
	}

	if err := json.Unmarshal(response, &resp); err != nil {
		logger.Error.Println(s.txID, " - couldn't bind Unmarshal in struct:", err)
		return resp, cod, err
	}

	return resp, cod, nil
}
