package reniec

import (
	"backend-ccff/internal/models"
	"backend-ccff/pkg/reniec/dni"
	"github.com/jmoiron/sqlx"
)

type ServerReniec struct {
	Dni dni.PortsServerReniec
}

func NewServerReniec(db *sqlx.DB, user *models.User, txID string) *ServerReniec {
	repoDni := dni.FactoryStorage(db, user, txID)
	return &ServerReniec{
		Dni: dni.NewReniecService(repoDni, user, txID),
	}
}
