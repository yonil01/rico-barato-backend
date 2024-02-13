package user_information_personal

import (
	"github.com/jmoiron/sqlx"

	"backend-comee/internal/logger"
	"backend-comee/internal/models"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesUserInformationPersonalRepository interface {
	create(m *UserInformationPersonal) error
	update(m *UserInformationPersonal) error
	delete(id int) error
	getByID(id int) (*UserInformationPersonal, error)
	getAll() ([]*UserInformationPersonal, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesUserInformationPersonalRepository {
	var s ServicesUserInformationPersonalRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return newUserInformationPersonalSqlServerRepository(db, user, txID)
	case Postgresql:
		return newUserInformationPersonalPsqlRepository(db, user, txID)
	case Oracle:
		return newUserInformationPersonalOrclRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
