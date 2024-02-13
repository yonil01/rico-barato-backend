package entity

import (
	"backend-comee/internal/models"
	"backend-comee/pkg/entity/food"
	"backend-comee/pkg/entity/information_entity"
	list_ListAttendance "backend-comee/pkg/entity/list-attendance"
	"backend-comee/pkg/entity/user_information_personal"
	"github.com/jmoiron/sqlx"
)

type ServerEntity struct {
	ListAttendance          list_ListAttendance.PortsServerListAttendance
	Food                    food.PortsServerFood
	UserInformationPersonal user_information_personal.PortsServerUserInformationPersonal
	InformationEntity       information_entity.PortsServerInformationEntity
}

func NewServerEntity(db *sqlx.DB, user *models.User, txID string) *ServerEntity {
	repoListAttendance := list_ListAttendance.FactoryStorage(db, user, txID)
	repoFood := food.FactoryStorage(db, user, txID)
	repoUserInformationPersonal := user_information_personal.FactoryStorage(db, user, txID)
	repoInformationEntity := information_entity.FactoryStorage(db, user, txID)
	return &ServerEntity{
		ListAttendance:          list_ListAttendance.NewListAttendanceService(repoListAttendance, user, txID),
		Food:                    food.NewFoodService(repoFood, user, txID),
		UserInformationPersonal: user_information_personal.NewUserInformationPersonalService(repoUserInformationPersonal, user, txID),
		InformationEntity:       information_entity.NewInformationEntityService(repoInformationEntity, user, txID),
	}
}
