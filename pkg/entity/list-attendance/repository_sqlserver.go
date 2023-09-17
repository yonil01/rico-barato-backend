package attendance

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

func newAttendanceSqlServerRepository(db *sqlx.DB, user *models.User, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *sqlserver) create(m *Attendance) error {
	var id int
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	const sqlInsert = `INSERT INTO entity.attendance_event_users (id_user, id_event, is_disable, is_delete, user_id, created_at, updated_at) VALUES (@id_user, @id_event, @is_disable, @is_delete, @user_id, @created_at, @updated_at) SELECT ID = convert(bigint, SCOPE_IDENTITY()) `
	stmt, err := s.DB.Prepare(sqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		sql.Named("id_user", m.IdUser),
		sql.Named("id_event", m.IdEvent),
		sql.Named("is_disable", m.IsDisable),
		sql.Named("is_delete", m.IsDelete),
		sql.Named("user_id", m.UserId),
		sql.Named("created_at", m.CreatedAt),
		sql.Named("updated_at", m.UpdatedAt),
	).Scan(&id)
	if err != nil {
		return err
	}
	m.ID = int(id)
	return nil
}

// Update actualiza un registro en la BD
func (s *sqlserver) update(m *Attendance) error {
	date := time.Now()
	m.UpdatedAt = date
	const sqlUpdate = `UPDATE entity.attendance SET id_user = :id_user, id_event = :id_event, is_disable = :is_disable, is_delete = :is_delete, user_id = :user_id, updated_at = :updated_at WHERE id = :id `
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
func (s *sqlserver) delete(id int) error {
	const sqlDelete = `DELETE FROM entity.attendance WHERE id = :id `
	m := Attendance{ID: id}
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
func (s *sqlserver) getByID(id int) (*Attendance, error) {
	const sqlGetByID = `SELECT convert(nvarchar(50), id) id , id_user, id_event, is_disable, is_delete, user_id, created_at, updated_at FROM entity.attendance  WITH (NOLOCK)  WHERE id = @id `
	mdl := Attendance{}
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
func (s *sqlserver) getAll() ([]*Attendance, error) {
	var ms []*Attendance
	const sqlGetAll = `SELECT convert(nvarchar(50), id) id , id_user, id_event, is_disable, is_delete, user_id, created_at, updated_at FROM entity.attendance  WITH (NOLOCK) `

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
