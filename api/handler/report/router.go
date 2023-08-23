package report

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterReport(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerEvents{dB: db, txID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	getWork := v1.Group("/report")
	getWork.Post("", h.GetInfoProcedure)
}
