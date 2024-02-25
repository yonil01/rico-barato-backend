package food

import (
	"backend-comee/internal/logger"
	"backend-comee/internal/models"
	"backend-comee/internal/msgs"
	"backend-comee/pkg/doc"
	"backend-comee/pkg/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
	"sync"
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

	var reqFoods []*models.Food

	for _, obj := range req {
		reqFile, err := srvDoc.Files.GetFilesByEntityId(obj.ID, 2)
		if err != nil {
			logger.Error.Printf("Couldn't insert suffragers: %v", err)
			res.Code, res.Type, res.Msg = msg.GetByCode(99)
			return c.Status(http.StatusAccepted).JSON(res)
		}

		rqEntity, cod, err := srvEntity.InformationEntity.GetInformationEntityByID(obj.EntityId)
		if err != nil {
			logger.Error.Printf("Couldn't insert suffragers: %v", err)
			res.Code, res.Type, res.Msg = msg.GetByCode(cod)
			return c.Status(http.StatusAccepted).JSON(res)
		}

		fileEntity, err := srvDoc.Files.GetFilesByEntityId(rqEntity.ID, 1)
		if err != nil {
			logger.Error.Printf("Couldn't insert suffragers: %v", err)
			res.Code, res.Type, res.Msg = msg.GetByCode(99)
			return c.Status(http.StatusAccepted).JSON(res)
		}
		rqEntity.File = fileEntity
		obj.Entity = rqEntity
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

	var reqFoods []*models.Food

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

func (h *handlerFoods) GetFoodsByEntityWithCoordinate(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseAllFood{Error: true}
	srvEntity := entity.NewServerEntity(h.dB, h.user, h.txID)
	srvDoc := doc.NewServerEntity(h.dB, h.user, h.txID)
	req := Coordinate{}

	err := c.BodyParser(&req)
	if err != nil {
		logger.Error.Printf("couldn't bind model BodyParser: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1)
		res.Msg = "couldn't bind model RequestMetadata"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	entities, err := srvEntity.InformationEntity.GetEntityByCoordinate(req.Long, req.Lat, req.Amount)
	if err != nil {
		logger.Error.Printf("Couldn't GetEntityByCoordinate: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(99)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	var dataFood []*models.Food
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, entity := range entities {
		wg.Add(1)
		go func(entity *models.Entity) {
			defer wg.Done()

			foods, err := srvEntity.Food.GetFoodsByEntityId(entity.ID)
			if err != nil {
				logger.Error.Printf("Couldn't GetFoodsByEntityId: %v", err)
				mu.Lock()
				res.Code, res.Type, res.Msg = msg.GetByCode(99)
				mu.Unlock()
				return
			}

			for _, food := range foods {
				file, err := srvDoc.Files.GetFilesByEntityId(food.ID, 2)
				if err != nil {
					logger.Error.Printf("Couldn't get Files: %v", err)
					mu.Lock()
					res.Code, res.Type, res.Msg = msg.GetByCode(99)
					mu.Unlock()
					return
				}
				food.File = file
				food.Entity = entity

				mu.Lock()
				dataFood = append(dataFood, food)
				mu.Unlock()
			}
		}(entity)
	}

	wg.Wait()

	res.Data = dataFood
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}
