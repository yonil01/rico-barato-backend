package comments

import (
	"backend-comee/internal/logger"
	"backend-comee/internal/models"
	"backend-comee/internal/msgs"
	"backend-comee/pkg/doc"
	"backend-comee/pkg/doc/comments"
	"backend-comee/pkg/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
)

type handlerInfoBasicEntitys struct {
	dB   *sqlx.DB
	user *models.User
	txID string
}

func (h *handlerInfoBasicEntitys) GetInfoBasicEntities(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseAllComments{Error: true}
	srvEntity := entity.NewServerEntity(h.dB, h.user, h.txID)

	_, err := srvEntity.InformationEntity.GetAllInformationEntity()
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(99)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = nil
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerInfoBasicEntitys) CreateComment(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseComments{Error: true}
	srvDoc := doc.NewServerEntity(h.dB, h.user, h.txID)
	srvEntity := entity.NewServerEntity(h.dB, h.user, h.txID)
	req := RequestComments{}

	err := c.BodyParser(&req)
	if err != nil {
		logger.Error.Printf("couldn't bind model BodyParser: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1)
		res.Msg = "couldn't bind model RequestMetadata"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	/*usr, err := middleware.GetUser(c)
	if err != nil {
		res.Error = true
		res.Code = 99
		res.Msg = "Error in token"
		return c.Status(http.StatusOK).JSON(res)
	}*/

	resData, cod, err := srvDoc.Comment.CreateComment(req.UserId, req.EntityId, req.Value, req.Start, 0)
	if err != nil {
		logger.Error.Printf("Couldn'tCreateComment: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(cod)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	usr, err := srvEntity.UserInformationPersonal.GetUserInformationPersonalByUserId(resData.UserId)
	if err != nil {
		logger.Error.Printf("Couldn't GetUserByID: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(cod)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	resData.User = usr

	res.Data = resData
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerInfoBasicEntitys) GetCommentsByEntityId(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseAllComments{Error: true}
	srvDoc := doc.NewServerEntity(h.dB, h.user, h.txID)
	srvEntity := entity.NewServerEntity(h.dB, h.user, h.txID)

	entityId := c.Params("id")
	intId, err := strconv.Atoi(entityId)
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(99)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	req, err := srvDoc.Comment.GetCommentByEntityID(intId)
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(99)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	var commentsAll []*comments.Comment
	for _, obj := range req {
		reqUser, err := srvEntity.UserInformationPersonal.GetUserInformationPersonalByUserId(obj.UserId)
		if err != nil {
			logger.Error.Printf("Couldn't GetUserInformationPersonalByUserId: %v", err)
			res.Code, res.Type, res.Msg = msg.GetByCode(99)
			return c.Status(http.StatusAccepted).JSON(res)
		}
		obj.User = reqUser
		commentsAll = append(commentsAll, obj)
	}

	res.Data = commentsAll
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}
