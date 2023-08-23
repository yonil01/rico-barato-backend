package attendance

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterAttendance(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerAttendance{dB: db, txID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	getWork := v1.Group("/attendance")
	getWork.Post("create", h.CreateAttendance)
}
