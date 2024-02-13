package list_ListAttendance

import (
	"backend-comee/internal/models"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

// psql estructura de conexión a la BD de postgresql
type psql struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func newListAttendancePsqlRepository(db *sqlx.DB, user *models.User, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *psql) create(m *ListAttendance) error {

	const psqlInsert = `INSERT INTO entity.ListAttendance (id_user, id_event, is_disable, is_delete, user_id) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at`
	stmt, err := s.DB.Prepare(psqlInsert)
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
func (s *psql) update(m *ListAttendance) error {
	date := time.Now()
	m.UpdatedAt = date
	const psqlUpdate = `UPDATE entity.ListAttendance SET id_user = :id_user, id_event = :id_event, is_disable = :is_disable, is_delete = :is_delete, user_id = :user_id, updated_at = :updated_at WHERE id = :id `
	rs, err := s.DB.NamedExec(psqlUpdate, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("Dev-cff:108")
	}
	return nil
}

// Delete elimina un registro de la BD
func (s *psql) delete(id int) error {
	const psqlDelete = `DELETE FROM entity.ListAttendance WHERE id = :id `
	m := ListAttendance{ID: id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("Dev-cff:108")
	}
	return nil
}

// GetByID consulta un registro por su ID
func (s *psql) getByID(id int) (*ListAttendance, error) {
	const psqlGetByID = `SELECT id , id_user, id_event, is_disable, is_delete, user_id, created_at, updated_at FROM entity.ListAttendance WHERE id = $1 `
	mdl := ListAttendance{}
	err := s.DB.Get(&mdl, psqlGetByID, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

// GetAll consulta todos los registros de la BD
func (s *psql) getAll() ([]*ListAttendance, error) {
	var ms []*ListAttendance
	const psqlGetAll = ` SELECT id , id_user, id_event, is_disable, is_delete, user_id, created_at, updated_at FROM entity.ListAttendance `

	err := s.DB.Select(&ms, psqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *psql) getListAttendanceUser() ([]*ListAttendance, error) {
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
