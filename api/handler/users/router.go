package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterUsers(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerUsers{dB: db, txID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	getWork := v1.Group("/users")
	getWork.Post("create", h.CreateUser)
	getWork.Get("all", h.GetUsers)
	getWork.Get("/:id", h.GetUserById)
	getWork.Get("delete/:id", h.DeleteUser)
}
