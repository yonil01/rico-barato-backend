package reniec

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/logger"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/models"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/msgs"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/pkg/reniec"
	"net/http"
)

type handlerReniec struct {
	dB   *sqlx.DB
	user *models.User
	txID string
}

func (h *handlerReniec) Reniec(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseReniec{Error: true}
	srvReniec := reniec.NewServerReniec(h.dB, h.user, h.txID)
	rqReniec := RequestReniec{}

	err := c.BodyParser(&rqReniec)
	if err != nil {
		logger.Error.Printf("couldn't bind model RequestMetadata: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1)
		res.Msg = "couldn't bind model RequestMetadata"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	req, cod, err := srvReniec.Dni.GetConsultReniecByDni(rqReniec.Dni)
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(cod)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = req
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}
