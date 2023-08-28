package Roles

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterRoles(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerRoles{dB: db, txID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	getWork := v1.Group("/roles")
	getWork.Get("", h.GetRoles)
	getWork.Post("create", h.CreateRoles)
	getWork.Post("update", h.UpdateRoles)
	getWork.Get("delete/:id", h.DeleteRole)
	getWork.Get(":id", h.GetRole)
	getWork.Post("name", h.GetRoleByName)

	roleUser := v1.Group("/role_user")
	roleUser.Get("", h.GetRolesUser)
	roleUser.Get("/:id", h.GetRoleUser)
	roleUser.Get("/delete/:id", h.DeleteRoleUser)
	roleUser.Get("all/:id", h.GetRoleUserByUser)
	roleUser.Post("create", h.CreateRoleUser)
}
