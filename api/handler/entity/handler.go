package entity

import (
	"backend-comee/internal/logger"
	"backend-comee/internal/models"
	"backend-comee/internal/msgs"
	"backend-comee/pkg/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type handlerInfoBasicEntitys struct {
	dB   *sqlx.DB
	user *models.User
	txID string
}

func (h *handlerInfoBasicEntitys) GetInfoBasicEntities(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseAllInfoEntity{Error: true}
	srvEntity := entity.NewServerEntity(h.dB, h.user, h.txID)

	req, err := srvEntity.InformationEntity.GetAllInformationEntity()
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(99)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = req
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerInfoBasicEntitys) CreateEntity(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseInfoEntity{Error: true}
	srvEntity := entity.NewServerEntity(h.dB, h.user, h.txID)
	req := RequestEntity{}

	err := c.BodyParser(&req)
	if err != nil {
		logger.Error.Printf("couldn't bind model BodyParser: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1)
		res.Msg = "couldn't bind model RequestMetadata"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	/*usr, err := middleware.GetUser(c)
	if err != nil {
		res.Error = true
		res.Code = 99
		res.Msg = "Error in token"
		return c.Status(http.StatusOK).JSON(res)
	}*/

	resData, cod, err := srvEntity.InformationEntity.CreateInformationEntity(req.UserEntityId, req.Name, req.Description, req.Telephone, req.Mobile, req.LocationX, req.LocationY, 0, 0, "33ac98cd-cac7-4eb7-8efe-d0e5264b4fd2")
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(cod)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = resData
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}
