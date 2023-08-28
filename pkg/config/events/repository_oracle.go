package events

import (
	"backend-ccff/internal/models"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

// Orcl estructura de conexión a la BD de Oracle
type orcl struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func newEventsOrclRepository(db *sqlx.DB, user *models.User, txID string) *orcl {
	return &orcl{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *orcl) create(m *Events) error {
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	const osqlInsert = `INSERT INTO cfg.events (id ,name, description, event_date, is_disable, is_delete, user_id, created_at, updated_at)  VALUES (:id ,:name, :description, :event_date, :is_disable, :is_delete, :user_id,:created_at, :updated_at) `
	rs, err := s.DB.NamedExec(osqlInsert, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("Dev-cff:108")
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *orcl) update(m *Events) error {
	date := time.Now()
	m.UpdatedAt = date
	const osqlUpdate = `UPDATE cfg.events SET name = :name, description = :description, event_date = :event_date, is_disable = :is_disable, is_delete = :is_delete, user_id = :user_id, updated_at = :updated_at WHERE id = :id  `
	rs, err := s.DB.NamedExec(osqlUpdate, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("Dev-cff:108")
	}
	return nil
}

// Delete elimina un registro de la BD
func (s *orcl) delete(id string) error {
	const osqlDelete = `DELETE FROM cfg.events WHERE id = :id `
	m := Events{ID: id}
	rs, err := s.DB.NamedExec(osqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("Dev-cff:108")
	}
	return nil
}

// GetByID consulta un registro por su ID
func (s *orcl) getByID(id string) (*Events, error) {
	const osqlGetByID = `SELECT id , name, description, event_date, is_disable, is_delete, user_id, created_at, updated_at FROM cfg.events WHERE id = :1 `
	mdl := Events{}
	err := s.DB.Get(&mdl, osqlGetByID, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

// GetAll consulta todos los registros de la BD
func (s *orcl) getAll() ([]*Events, error) {
	var ms []*Events
	const osqlGetAll = ` SELECT id , name, description, event_date, is_disable, is_delete, user_id, created_at, updated_at FROM cfg.events `

	err := s.DB.Select(&ms, osqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
