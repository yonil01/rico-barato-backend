package user_entity

import (
	"backend-comee/internal/logger"
	"backend-comee/internal/models"
	"backend-comee/internal/msgs"
	"backend-comee/pkg/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type handlerInfoBasicPersons struct {
	dB   *sqlx.DB
	user *models.User
	txID string
}

func (h *handlerInfoBasicPersons) GetUserEntity(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseAllInfoBasicPerson{Error: true}
	srvAuth := auth.NewServerAuth(h.dB, h.user, h.txID)

	req, err := srvAuth.UserEntity.GetAllUserEntity()
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(99)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = req
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerInfoBasicPersons) CreateUserEntity(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseInfoBasicPerson{Error: true}
	srvAuth := auth.NewServerAuth(h.dB, h.user, h.txID)
	req := RequestUserEntity{}

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

	resData, cod, err := srvAuth.UserEntity.CreateUserEntity(req.Id, req.Dni, req.Name, req.Lastname, req.Email, req.Password, 0, 0, "33ac98cd-cac7-4eb7-8efe-d0e5264b4fd2")
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(cod)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = resData
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerInfoBasicPersons) Login(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseInfoBasicPerson{Error: true}
	srvAuth := auth.NewServerAuth(h.dB, h.user, h.txID)
	req := RequestUserEntity{}

	err := c.BodyParser(&req)
	if err != nil {
		logger.Error.Printf("couldn't bind model BodyParser: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1)
		res.Msg = "couldn't bind model RequestMetadata"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	resData, cod, err := srvAuth.UserEntity.CreateUserEntity(req.Id, req.Dni, req.Name, req.Lastname, req.Email, req.Password, 0, 0, "33ac98cd-cac7-4eb7-8efe-d0e5264b4fd2")
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(cod)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = resData
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}
