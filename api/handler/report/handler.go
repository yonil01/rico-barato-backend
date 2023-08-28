package report

import (
	"backend-ccff/internal/logger"
	"backend-ccff/internal/middleware"
	"backend-ccff/internal/models"
	"backend-ccff/internal/msgs"
	"backend-ccff/pkg/transactions/report"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type handlerEvents struct {
	dB   *sqlx.DB
	user *models.User
	txID string
}

func (h *handlerEvents) GetInfoProcedure(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseReport{Error: true}
	rqReport := RequestReport{}

	err := c.BodyParser(&rqReport)
	if err != nil {
		logger.Error.Printf("couldn't bind model BodyParser: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1)
		res.Msg = "couldn't bind model RequestMetadata"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	usr, err := middleware.GetUser(c)
	if err != nil {
		res.Error = true
		res.Code = 99
		res.Msg = "Error in token"
		return c.Status(http.StatusOK).JSON(res)
	}

	repoReport := report.FactoryStorage(h.dB, usr, h.txID)
	srvReport := report.NewReportService(repoReport, usr, h.txID)

	rs, err := srvReport.ExecuteReport(rqReport.Procedure, rqReport.Parameters.(map[string]interface{}), 1)
	if err != nil {
		logger.Error.Println("couldn't ExecuteReport")

		codeInt, typeString, message := msg.GetByCode(22)
		res.Code = codeInt
		res.Type = typeString
		res.Msg = message
		return c.Status(http.StatusOK).JSON(res)
	}
	res.Data = rs
	codeInt, typeString, message := msg.GetByCode(29)
	res.Code = codeInt
	res.Type = typeString
	res.Msg = message
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}
