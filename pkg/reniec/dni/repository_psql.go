package dni

import (
	"backend-ccff/internal/models"
	"github.com/jmoiron/sqlx"
)

type psql struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func newReniecPsqlRepository(db *sqlx.DB, user *models.User, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}
