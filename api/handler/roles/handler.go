package Roles

import (
	"backend-ccff/internal/logger"
	"backend-ccff/internal/middleware"
	"backend-ccff/internal/models"
	"backend-ccff/internal/msgs"
	"backend-ccff/pkg/auth"
	"backend-ccff/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type handlerRoles struct {
	dB   *sqlx.DB
	user *models.User
	txID string
}

func (h *handlerRoles) CreateRoles(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseRoles{Error: true}

	rqARoles := RequestRoles{}

	err := c.BodyParser(&rqARoles)
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

	srvAuth := auth.NewServerAuth(h.dB, usr, h.txID)

	req, cod, err := srvAuth.Roles.CreateRole(rqARoles.Id, rqARoles.Name, rqARoles.Description, 1, true)
	if err != nil {
		logger.Error.Printf("Couldn't insert CreateRoles: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(cod)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = req
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerRoles) UpdateRoles(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseRoles{Error: true}

	rqARoles := RequestRoles{}

	err := c.BodyParser(&rqARoles)
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

	srvAuth := auth.NewServerAuth(h.dB, usr, h.txID)

	req, cod, err := srvAuth.Roles.UpdateRole(rqARoles.Id, rqARoles.Name, rqARoles.Description, 1, true)
	if err != nil {
		logger.Error.Printf("Couldn't insert CreateRoles: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(cod)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = req
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerRoles) GetRoles(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseAllRoles{Error: true}
	srvAuth := auth.NewServerAuth(h.dB, h.user, h.txID)

	req, err := srvAuth.Roles.GetAllRole()
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(99)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = req
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerRoles) DeleteRole(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseRoles{Error: true}
	usr, err := middleware.GetUser(c)
	if err != nil {
		res.Error = true
		res.Code = 99
		res.Msg = "Error in token"
		return c.Status(http.StatusOK).JSON(res)
	}
	roleId := c.Params("id")
	srvRole := auth.NewServerAuth(h.dB, usr, h.txID)

	cod, err := srvRole.Roles.DeleteRole(roleId)
	if err != nil {
		logger.Error.Printf("Couldn't  GetModulesRole: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(cod)
		return c.Status(http.StatusAccepted).JSON(res)
	}
	res.Data = nil
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerRoles) GetRole(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseRoles{Error: true}
	usr, err := middleware.GetUser(c)
	if err != nil {
		res.Error = true
		res.Code = 99
		res.Msg = "Error in token"
		return c.Status(http.StatusOK).JSON(res)
	}
	roleId := c.Params("id")
	srvRole := auth.NewServerAuth(h.dB, usr, h.txID)

	resp, cod, err := srvRole.Roles.GetRoleByID(roleId)
	if err != nil {
		logger.Error.Printf("Couldn't  GetModulesRole: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(cod)
		return c.Status(http.StatusAccepted).JSON(res)
	}
	res.Data = resp
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerRoles) GetRoleByName(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseRoles{Error: true}

	rqARoles := RequestRoleName{}

	err := c.BodyParser(&rqARoles)
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

	srvAuth := auth.NewServerAuth(h.dB, usr, h.txID)

	req, cod, err := srvAuth.Roles.GetRoleByName(rqARoles.Name)
	if err != nil {
		logger.Error.Printf("Couldn't insert GetRoleByName: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(cod)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = req
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerRoles) GetRolesUser(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseAllRolesUser{Error: true}
	srvConfig := config.NewServerConfig(h.dB, h.user, h.txID)

	req, err := srvConfig.RoleUser.GetAllRoleUser()
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(99)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = req
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerRoles) GetRoleUser(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseRolesUser{Error: true}
	usr, err := middleware.GetUser(c)
	if err != nil {
		res.Error = true
		res.Code = 99
		res.Msg = "Error in token"
		return c.Status(http.StatusOK).JSON(res)
	}
	roleId := c.Params("id")
	srvRoleUser := config.NewServerConfig(h.dB, usr, h.txID)

	resp, cod, err := srvRoleUser.RoleUser.GetRoleUserByID(roleId)
	if err != nil {
		logger.Error.Printf("Couldn't  GetModulesRole: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(cod)
		return c.Status(http.StatusAccepted).JSON(res)
	}
	res.Data = resp
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerRoles) GetRoleUserByUser(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseAllRolesUser{Error: true}
	usr, err := middleware.GetUser(c)
	if err != nil {
		res.Error = true
		res.Code = 99
		res.Msg = "Error in token"
		return c.Status(http.StatusOK).JSON(res)
	}
	roleId := c.Params("id")
	srvRoleUser := config.NewServerConfig(h.dB, usr, h.txID)

	resp, err := srvRoleUser.RoleUser.GetAllRoleUserByUser(roleId)
	if err != nil {
		logger.Error.Printf("Couldn't  GetModulesRole: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(99)
		return c.Status(http.StatusAccepted).JSON(res)
	}
	res.Data = resp
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerRoles) DeleteRoleUser(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseRolesUser{Error: true}
	usr, err := middleware.GetUser(c)
	if err != nil {
		res.Error = true
		res.Code = 99
		res.Msg = "Error in token"
		return c.Status(http.StatusOK).JSON(res)
	}
	roleId := c.Params("id")
	srvRoleUser := config.NewServerConfig(h.dB, usr, h.txID)

	cod, err := srvRoleUser.RoleUser.DeleteRoleUser(roleId)
	if err != nil {
		logger.Error.Printf("Couldn't  GetModulesRole: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(cod)
		return c.Status(http.StatusAccepted).JSON(res)
	}
	res.Data = nil
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerRoles) CreateRoleUser(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseRolesUser{Error: true}

	rqARoles := RequestRoleUser{}

	err := c.BodyParser(&rqARoles)
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

	srvConfig := config.NewServerConfig(h.dB, usr, h.txID)

	req, cod, err := srvConfig.RoleUser.CreateRoleUser(rqARoles.Id, rqARoles.UserId, rqARoles.RoleId)
	if err != nil {
		logger.Error.Printf("Couldn't insert CreateRoles: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(cod)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = req
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}
