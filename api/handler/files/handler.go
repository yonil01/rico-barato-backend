package files

import (
	"backend-comee/internal/logger"
	"backend-comee/internal/models"
	"backend-comee/internal/msgs"
	"backend-comee/pkg/doc"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type handlerFiless struct {
	dB   *sqlx.DB
	user *models.User
	txID string
}

func (h *handlerFiless) GetFilesByEntityId(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseAllFiles{Error: true}
	srvPerson := doc.NewServerEntity(h.dB, h.user, h.txID)

	req, err := srvPerson.Files.GetAllFiles()
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(99)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = req
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerFiless) CreateFiles(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseFiles{Error: true}
	srvDoc := doc.NewServerEntity(h.dB, h.user, h.txID)
	req := RequestFiles{}

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

	_, cod, err := srvDoc.Files.CreateFiles(req.EntityId, req.Path, req.TypeDocument, req.TypeEntity, "33ac98cd-cac7-4eb7-8efe-d0e5264b4fd2", 0)
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(cod)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = &req
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}
