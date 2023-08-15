package upload_metadata

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/generator"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/logger"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/models"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/msgs"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/pkg/indra"
	"net/http"
)

type handlerUploadMetadata struct {
	dB   *sqlx.DB
	user *models.User
	txID string
}

func (h *handlerUploadMetadata) uploadMetadata(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseUploadMetadata{Error: true}
	srvIndra := indra.NewServerIndra(h.dB, h.user, h.txID)
	rqUploadMetadata := RequestProcess{}

	err := c.BodyParser(&rqUploadMetadata)
	if err != nil {
		logger.Error.Printf("couldn't bind model RequestMetadata: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1)
		res.Msg = "couldn't bind model RequestMetadata"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	for i, metadatum := range rqUploadMetadata.Metadata {
		fmt.Print(i)
		fmt.Print(metadatum)
	}

	cod, err := srvIndra.SrvUploadMetadata.UpdateMetadata(rqUploadMetadata.Metadata)
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(cod)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerUploadMetadata) generateMetadata(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseUploadMetadata{Error: true}
	srvIndra := indra.NewServerIndra(h.dB, h.user, h.txID)
	rqUploadMetadata := RequestGenerate{}

	err := c.BodyParser(&rqUploadMetadata)
	if err != nil {
		logger.Error.Printf("couldn't bind model RequestMetadata: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1)
		res.Msg = "couldn't bind model RequestMetadata"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	value, err := srvIndra.SrvUploadMetadata.GetIdsAutofillValueByEntityAttributeAndValue(rqUploadMetadata.TypeInput, rqUploadMetadata.InputData)
	if err != nil {
		logger.Error.Printf("Couldn't GetIdsAutofillValueByEntityAttributeAndValue: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(99)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	pathFile, err := generator.GeneratorXLSX(value)
	if err != nil {
		logger.Error.Printf("Couldn't GeneratorXLSX: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(99)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	res.Data = pathFile
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}
