package list_ListAttendance

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

func newListAttendanceOrclRepository(db *sqlx.DB, user *models.User, txID string) *orcl {
	return &orcl{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *orcl) create(m *ListAttendance) error {
	const osqlInsert = `INSERT INTO entity.ListAttendance (id_user, id_event, is_disable, is_delete, user_id, created_at, updated_at)  VALUES (:id_user, :id_event, :is_disable, :is_delete, :user_id, sysdate, sysdate) RETURNING id into id, created_at into created_at, updated_at into updated_at `
	stmt, err := s.DB.Prepare(osqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		m.IsDelete,
	).Scan(&m.ID, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *orcl) update(m *ListAttendance) error {
	date := time.Now()
	m.UpdatedAt = date
	const osqlUpdate = `UPDATE entity.ListAttendance SET id_user = :id_user, id_event = :id_event, is_disable = :is_disable, is_delete = :is_delete, user_id = :user_id, updated_at = :updated_at WHERE id = :id  `
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
	const osqlDelete = `DELETE FROM entity.ListAttendance WHERE id = :id `
	m := ListAttendance{ID: id}
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
func (s *orcl) getByID(id int) (*ListAttendance, error) {
	const osqlGetByID = `SELECT id , id_user, id_event, is_disable, is_delete, user_id, created_at, updated_at FROM entity.ListAttendance WHERE id = :1 `
	mdl := ListAttendance{}
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
func (s *orcl) getAll() ([]*ListAttendance, error) {
	var ms []*ListAttendance
	const osqlGetAll = ` SELECT id , id_user, id_event, is_disable, is_delete, user_id, created_at, updated_at FROM entity.ListAttendance `

	err := s.DB.Select(&ms, osqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *orcl) getListAttendanceUser() ([]*ListAttendance, error) {
	userId := s.user.ID
	const sqlGetByID = `SELECT ev.description, ae.is_delete, us.names, ae.created_at, ae.updated_at
FROM [entity].[attendance_event_users]   ae
join [cfg].[events] ev on ev.id = ae.id_event 
join auth.users us on us.id = ae.user_id
WHERE id_user = @id_user `
	var mdl []*ListAttendance
	err := s.DB.Select(&mdl, sqlGetByID, sql.Named("id_user", userId))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return mdl, err
	}
	return mdl, nil
}
