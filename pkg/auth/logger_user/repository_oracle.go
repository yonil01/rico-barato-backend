package logger_user

import (
	"database/sql"
	"fmt"
	"time"

	"backend-comee/internal/models"
	"github.com/jmoiron/sqlx"
)

// Orcl estructura de conexión a la BD de Oracle
type orcl struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func newLoggerUserOrclRepository(db *sqlx.DB, user *models.User, txID string) *orcl {
	return &orcl{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *orcl) create(m *LoggerUser) error {
	const osqlInsert = `INSERT INTO wf.logger_user (user_id, ip, event, created_at, updated_at)  VALUES (:user_id, :ip, :event, sysdate, sysdate) RETURNING id into id, created_at into created_at, updated_at into updated_at `
	stmt, err := s.DB.Prepare(osqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		m.UserId,
		m.Ip,
		m.Event,
	).Scan(&m.ID, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *orcl) update(m *LoggerUser) error {
	date := time.Now()
	m.UpdatedAt = date
	const osqlUpdate = `UPDATE wf.logger_user SET user_id = :user_id, ip = :ip, event = :event, updated_at = :updated_at WHERE id = :id  `
	rs, err := s.DB.NamedExec(osqlUpdate, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// Delete elimina un registro de la BD
func (s *orcl) delete(id int) error {
	const osqlDelete = `DELETE FROM wf.logger_user WHERE id = :id `
	m := LoggerUser{ID: id}
	rs, err := s.DB.NamedExec(osqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// GetByID consulta un registro por su ID
func (s *orcl) getByID(id int) (*LoggerUser, error) {
	const osqlGetByID = `SELECT id , user_id, ip, event, created_at, updated_at FROM wf.logger_user WHERE id = :1 `
	mdl := LoggerUser{}
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
func (s *orcl) getAll() ([]*LoggerUser, error) {
	var ms []*LoggerUser
	const osqlGetAll = ` SELECT id , user_id, ip, event, created_at, updated_at FROM wf.logger_user `

	err := s.DB.Select(&ms, osqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
