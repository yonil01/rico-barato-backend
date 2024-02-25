package food

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterFood(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerFoods{dB: db, txID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	getWork := v1.Group("/food")
	getWork.Get("", h.GetFoods)
	getWork.Get("entity-id/:id/:type", h.GetFoodsByEntityId)
	getWork.Post("entity/coordinate", h.GetFoodsByEntityWithCoordinate)
	getWork.Post("create", h.CreateFood)
}
