package modules

import (
	"backend-ccff/internal/logger"
	"backend-ccff/internal/middleware"
	"backend-ccff/internal/models"
	"backend-ccff/internal/msgs"
	"backend-ccff/pkg/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type handlerModules struct {
	dB   *sqlx.DB
	user *models.User
	txID string
}

func (h *handlerModules) GetModulesByUser(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseModules{Error: true}

	rqModulesUser := RequestModulesUser{}

	usr, err := middleware.GetUser(c)
	if err != nil {
		res.Error = true
		res.Code = 99
		res.Msg = "Error in token"
		return c.Status(http.StatusOK).JSON(res)
	}

	srvModules := auth.NewServerAuth(h.dB, usr, h.txID)

	err = c.BodyParser(&rqModulesUser)
	if err != nil {
		logger.Error.Printf("couldn't bind model rqModulesUser: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1)
		res.Msg = "couldn't bind model RequestMetadata"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	req, err := srvModules.Modules.GetModulesByRoles(rqModulesUser.Ids, rqModulesUser.Ids, rqModulesUser.Type)
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(99)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = req
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerModules) GetModules(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseModules{Error: true}
	usr, err := middleware.GetUser(c)
	if err != nil {
		res.Error = true
		res.Code = 99
		res.Msg = "Error in token"
		return c.Status(http.StatusOK).JSON(res)
	}

	srvModules := auth.NewServerAuth(h.dB, usr, h.txID)

	req, err := srvModules.Modules.GetAllModule()
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(99)
		return c.Status(http.StatusAccepted).JSON(res)
	}
	res.Data = req
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerModules) GetModulesRoles(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseModulesRole{Error: true}
	usr, err := middleware.GetUser(c)
	if err != nil {
		res.Error = true
		res.Code = 99
		res.Msg = "Error in token"
		return c.Status(http.StatusOK).JSON(res)
	}
	roleId := c.Params("id")
	srvModules := auth.NewServerAuth(h.dB, usr, h.txID)

	req, err := srvModules.Modules.GetModulesRole(roleId)
	if err != nil {
		logger.Error.Printf("Couldn't  GetModulesRole: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(99)
		return c.Status(http.StatusAccepted).JSON(res)
	}
	res.Data = req
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerModules) DeleteModulesRoles(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseModulesRole{Error: true}
	usr, err := middleware.GetUser(c)
	if err != nil {
		res.Error = true
		res.Code = 99
		res.Msg = "Error in token"
		return c.Status(http.StatusOK).JSON(res)
	}
	id := c.Params("id")
	srvModules := auth.NewServerAuth(h.dB, usr, h.txID)

	cod, err := srvModules.Modules.DeleteModulesRole(id)
	if err != nil {
		logger.Error.Printf("Couldn't  GetModulesRole: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(cod)
		return c.Status(http.StatusAccepted).JSON(res)
	}
	res.Data = nil
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerModules) CreateModulesRoles(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseModuleRole{Error: true}
	rqEvent := RequestModulesRole{}
	usr, err := middleware.GetUser(c)
	if err != nil {
		res.Error = true
		res.Code = 99
		res.Msg = "Error in token"
		return c.Status(http.StatusOK).JSON(res)
	}
	err = c.BodyParser(&rqEvent)
	if err != nil {
		logger.Error.Printf("couldn't bind model BodyParser: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1)
		res.Msg = "couldn't bind model RequestMetadata"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	srvModules := auth.NewServerAuth(h.dB, usr, h.txID)

	resp, cod, err := srvModules.Modules.CreateModulesRole(rqEvent.Id, rqEvent.RoleId, rqEvent.ElementId)
	if err != nil {
		logger.Error.Printf("Couldn't  GetModulesRole: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(cod)
		return c.Status(http.StatusAccepted).JSON(res)
	}
	res.Data = resp
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}
