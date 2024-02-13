package user_information_personal

import (
	"backend-comee/internal/logger"
	"backend-comee/internal/models"
	"backend-comee/internal/msgs"
	"backend-comee/pkg/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type handlerUserInformationPersonal struct {
	dB   *sqlx.DB
	user *models.User
	txID string
}

func (h *handlerUserInformationPersonal) CreateUserInformationPersonal(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseUserInformationPersonal{Error: true}
	srvUserInformationPersonal := entity.NewServerEntity(h.dB, h.user, h.txID)
	rq := RequestUserInformationPersonal{}

	err := c.BodyParser(&rq)
	if err != nil {
		logger.Error.Printf("couldn't bind model RequestMetadata: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1)
		res.Msg = "couldn't bind model RequestMetadata"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	req, cod, err := srvUserInformationPersonal.UserInformationPersonal.CreateUserInformationPersonal(rq.UserId, rq.Gender, rq.Age)
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(cod)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = req
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerUserInformationPersonal) GetUserInformationPersonal(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseUserInformationPersonalAll{Error: true}
	srvUserInformationPersonal := entity.NewServerEntity(h.dB, h.user, h.txID)

	req, err := srvUserInformationPersonal.UserInformationPersonal.GetAllUserInformationPersonal()
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(99)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = req
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerUserInformationPersonal) DeleteUser(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseUserInformationPersonalAll{Error: true}
	srvUserInformationPersonal := entity.NewServerEntity(h.dB, h.user, h.txID)

	cod, err := srvUserInformationPersonal.UserInformationPersonal.DeleteUserInformationPersonal(1)
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(cod)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = nil
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerUserInformationPersonal) GetUserById(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseUserInformationPersonal{Error: true}
	srvUserInformationPersonal := entity.NewServerEntity(h.dB, h.user, h.txID)

	usr, cod, err := srvUserInformationPersonal.UserInformationPersonal.GetUserInformationPersonalByID(1)
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(cod)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = usr
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}
