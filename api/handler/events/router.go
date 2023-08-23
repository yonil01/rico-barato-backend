package events

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterEvents(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerEvents{dB: db, txID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	getWork := v1.Group("/event")
	getWork.Get("all", h.GetEvents)
	getWork.Post("create", h.CreateEvent)
}
