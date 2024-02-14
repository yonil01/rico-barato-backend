package user_entity

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterUserEntity(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerInfoBasicPersons{dB: db, txID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	getWork := v1.Group("/user-entity")
	getWork.Get("", h.GetUserEntity)
	getWork.Post("create", h.CreateUserEntity)
	getWork.Post("login", h.Login)
}
