package events

import (
	"backend-ccff/internal/models"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

// sqlServer estructura de conexión a la BD de mssql
type sqlserver struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func newEventsSqlServerRepository(db *sqlx.DB, user *models.User, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *sqlserver) create(m *Events) error {
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	const sqlInsert = `INSERT INTO cfg.events (id ,name, description, event_date, is_disable, is_delete, user_id, created_at, updated_at) VALUES (:id ,:name, :description, :event_date, :is_disable, :is_delete, :user_id, :created_at, :updated_at) `
	rs, err := s.DB.NamedExec(sqlInsert, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("Dev-cff:108")
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *sqlserver) update(m *Events) error {
	date := time.Now()
	m.UpdatedAt = date
	const sqlUpdate = `UPDATE cfg.events SET name = :name, description = :description, event_date = :event_date, is_disable = :is_disable, is_delete = :is_delete, user_id = :user_id, updated_at = :updated_at WHERE id = :id `
	rs, err := s.DB.NamedExec(sqlUpdate, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("Dev-cff:108")
	}
	return nil
}

// Delete elimina un registro de la BD
func (s *sqlserver) delete(id string) error {
	const sqlDelete = `DELETE FROM cfg.events WHERE id = :id `
	m := Events{ID: id}
	rs, err := s.DB.NamedExec(sqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("Dev-cff:108")
	}
	return nil
}

// GetByID consulta un registro por su ID
func (s *sqlserver) getByID(id string) (*Events, error) {
	const sqlGetByID = `SELECT convert(nvarchar(50), id) id , name, description, event_date, is_disable, is_delete, user_id, created_at, updated_at FROM cfg.events  WITH (NOLOCK)  WHERE id = @id `
	mdl := Events{}
	err := s.DB.Get(&mdl, sqlGetByID, sql.Named("id", id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

// GetAll consulta todos los registros de la BD
func (s *sqlserver) getAll() ([]*Events, error) {
	var ms []*Events
	const sqlGetAll = `SELECT convert(nvarchar(50), id) id , name, description, event_date, is_disable, is_delete, user_id, created_at, updated_at FROM cfg.events  WITH (NOLOCK) `

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
