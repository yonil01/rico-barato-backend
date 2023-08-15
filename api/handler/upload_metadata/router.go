package upload_metadata

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterUploadMetadata(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerUploadMetadata{dB: db, txID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	getWork := v1.Group("/upload-metadata")
	getWork.Post("", h.uploadMetadata)
	getWork.Post("generate", h.generateMetadata)

}
