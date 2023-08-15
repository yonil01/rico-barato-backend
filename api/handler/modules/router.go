package modules

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterModules(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerModules{dB: db, txID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	getWork := v1.Group("/modules")
	getWork.Post("user", h.GetModulesUser)
}
