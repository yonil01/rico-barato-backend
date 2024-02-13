package food

import (
	"backend-comee/internal/logger"
	"backend-comee/internal/models"
	"backend-comee/internal/msgs"
	"backend-comee/pkg/doc"
	"backend-comee/pkg/entity"
	"backend-comee/pkg/entity/food"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
)

type handlerFoods struct {
	dB   *sqlx.DB
	user *models.User
	txID string
}

func (h *handlerFoods) GetFoods(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseAllFood{Error: true}
	srvEntity := entity.NewServerEntity(h.dB, h.user, h.txID)

	req, err := srvEntity.Food.GetAllFood()
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(99)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	srvDoc := doc.NewServerEntity(h.dB, h.user, h.txID)

	var reqFoods []*food.Food

	for _, obj := range req {
		reqFile, err := srvDoc.Files.GetFilesByEntityId(obj.ID, 2)
		if err != nil {
			logger.Error.Printf("Couldn't insert suffragers: %v", err)
			res.Code, res.Type, res.Msg = msg.GetByCode(99)
			return c.Status(http.StatusAccepted).JSON(res)
		}
		obj.File = reqFile
		reqFoods = append(reqFoods, obj)
	}

	res.Data = req
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerFoods) CreateFood(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseFood{Error: true}
	srvEntity := entity.NewServerEntity(h.dB, h.user, h.txID)
	req := RequestFood{}

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

	resData, cod, err := srvEntity.Food.CreateFood(req.EntityId, req.Name, req.Description, req.Price,
		0, 0, 0, "33ac98cd-cac7-4eb7-8efe-d0e5264b4fd2")
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(cod)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = resData
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerFoods) GetFoodsByEntityId(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseAllFood{Error: true}
	srvEntity := entity.NewServerEntity(h.dB, h.user, h.txID)

	entityId := c.Params("id")
	intId, err := strconv.Atoi(entityId)
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(99)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	req, err := srvEntity.Food.GetFoodsByEntityId(intId)
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(99)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	entityType := c.Params("type")
	typeEntity, err := strconv.Atoi(entityType)
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(99)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	srvDoc := doc.NewServerEntity(h.dB, h.user, h.txID)

	var reqFoods []*food.Food

	for _, obj := range req {
		reqFile, err := srvDoc.Files.GetFilesByEntityId(obj.ID, typeEntity)
		if err != nil {
			logger.Error.Printf("Couldn't insert suffragers: %v", err)
			res.Code, res.Type, res.Msg = msg.GetByCode(99)
			return c.Status(http.StatusAccepted).JSON(res)
		}
		obj.File = reqFile
		reqFoods = append(reqFoods, obj)
	}

	res.Data = reqFoods
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}
