package attendance

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

func newAttendanceOrclRepository(db *sqlx.DB, user *models.User, txID string) *orcl {
	return &orcl{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *orcl) create(m *Attendance) error {
	const osqlInsert = `INSERT INTO entity.attendance (id_user, id_event, is_disable, is_delete, user_id, created_at, updated_at)  VALUES (:id_user, :id_event, :is_disable, :is_delete, :user_id, sysdate, sysdate) RETURNING id into id, created_at into created_at, updated_at into updated_at `
	stmt, err := s.DB.Prepare(osqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		m.IdUser,
		m.IdEvent,
		m.IsDisable,
		m.IsDelete,
		m.UserId,
	).Scan(&m.ID, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *orcl) update(m *Attendance) error {
	date := time.Now()
	m.UpdatedAt = date
	const osqlUpdate = `UPDATE entity.attendance SET id_user = :id_user, id_event = :id_event, is_disable = :is_disable, is_delete = :is_delete, user_id = :user_id, updated_at = :updated_at WHERE id = :id  `
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
func (s *orcl) delete(id int) error {
	const osqlDelete = `DELETE FROM entity.attendance WHERE id = :id `
	m := Attendance{ID: id}
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
func (s *orcl) getByID(id int) (*Attendance, error) {
	const osqlGetByID = `SELECT id , id_user, id_event, is_disable, is_delete, user_id, created_at, updated_at FROM entity.attendance WHERE id = :1 `
	mdl := Attendance{}
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
func (s *orcl) getAll() ([]*Attendance, error) {
	var ms []*Attendance
	const osqlGetAll = ` SELECT id , id_user, id_event, is_disable, is_delete, user_id, created_at, updated_at FROM entity.attendance `

	err := s.DB.Select(&ms, osqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
