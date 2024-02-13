package user_information_personal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterUserInformationPersonal(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerUserInformationPersonal{dB: db, txID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	getWork := v1.Group("/user-information-personal")
	getWork.Post("create", h.CreateUserInformationPersonal)
	getWork.Get("all", h.GetUserInformationPersonal)
	getWork.Get("/:id", h.GetUserById)
	getWork.Get("delete/:id", h.DeleteUser)
}
