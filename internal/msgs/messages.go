package msgs

import (
	"github.com/jmoiron/sqlx"
)

type Model struct {
	db *sqlx.DB
}

func (m *Model) GetByCode(code int) (int, string, string) {
	return code, "-", "-"
}
