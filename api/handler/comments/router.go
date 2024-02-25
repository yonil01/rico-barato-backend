package comments

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterComments(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerInfoBasicEntitys{dB: db, txID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	getWork := v1.Group("/comments")
	getWork.Get("", h.GetInfoBasicEntities)
	getWork.Post("create", h.CreateComment)
	getWork.Get("entity/:id", h.GetCommentsByEntityId)
}
