package reniec

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterReniec(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerReniec{dB: db, txID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	getWork := v1.Group("/reniec")
	getWork.Post("dni", h.Reniec)
}
