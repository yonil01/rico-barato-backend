package events

import (
	"backend-ccff/internal/logger"
	"backend-ccff/internal/middleware"
	"backend-ccff/internal/models"
	"backend-ccff/internal/msgs"
	"backend-ccff/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type handlerEvents struct {
	dB   *sqlx.DB
	user *models.User
	txID string
}

func (h *handlerEvents) GetEvents(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseAllEvent{Error: true}
	srvConfig := config.NewServerConfig(h.dB, h.user, h.txID)

	req, err := srvConfig.Event.GetAllEvents()
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(99)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = req
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerEvents) CreateEvent(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseEvent{Error: true}
	srvConfig := config.NewServerConfig(h.dB, h.user, h.txID)
	rqEvent := RequestEvent{}

	err := c.BodyParser(&rqEvent)
	if err != nil {
		logger.Error.Printf("couldn't bind model BodyParser: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1)
		res.Msg = "couldn't bind model RequestMetadata"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	usr, err := middleware.GetUser(c)
	if err != nil {
		res.Error = true
		res.Code = 99
		res.Msg = "Error in token"
		return c.Status(http.StatusOK).JSON(res)
	}

	req, cod, err := srvConfig.Event.CreateEvents(rqEvent.Id, rqEvent.Name, rqEvent.Description, rqEvent.EventDate, 0, 0, usr.ID)
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(cod)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = req
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}
