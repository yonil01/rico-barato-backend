package files

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterFiles(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerFiless{dB: db, txID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	getWork := v1.Group("/files")
	getWork.Get("", h.GetFilesByEntityId)
	getWork.Post("create", h.CreateFiles)
}
