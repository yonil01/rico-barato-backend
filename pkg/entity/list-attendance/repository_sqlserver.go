package list_ListAttendance

import (
	"backend-comee/internal/models"
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

func newListAttendanceSqlServerRepository(db *sqlx.DB, user *models.User, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *sqlserver) create(m *ListAttendance) error {
	var id int
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	const sqlInsert = `INSERT INTO entity.ListAttendance_event_users (id_user, id_event, is_disable, is_delete, user_id, created_at, updated_at) VALUES (@id_user, @id_event, @is_disable, @is_delete, @user_id, @created_at, @updated_at) SELECT ID = convert(bigint, SCOPE_IDENTITY()) `
	stmt, err := s.DB.Prepare(sqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		sql.Named("is_delete", m.IsDelete),
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
func (s *sqlserver) update(m *ListAttendance) error {
	date := time.Now()
	m.UpdatedAt = date
	const sqlUpdate = `UPDATE entity.ListAttendance SET id_user = :id_user, id_event = :id_event, is_disable = :is_disable, is_delete = :is_delete, user_id = :user_id, updated_at = :updated_at WHERE id = :id `
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
	const sqlDelete = `DELETE FROM entity.ListAttendance WHERE id = :id `
	m := ListAttendance{ID: id}
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
func (s *sqlserver) getByID(id int) (*ListAttendance, error) {
	const sqlGetByID = `SELECT convert(nvarchar(50), id) id , id_user, id_event, is_disable, is_delete, user_id, created_at, updated_at FROM entity.ListAttendance  WITH (NOLOCK)  WHERE id = @id `
	mdl := ListAttendance{}
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
func (s *sqlserver) getAll() ([]*ListAttendance, error) {
	var ms []*ListAttendance
	const sqlGetAll = `SELECT convert(nvarchar(50), id) id , id_user, id_event, is_disable, is_delete, user_id, created_at, updated_at FROM entity.ListAttendance  WITH (NOLOCK) `

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *sqlserver) getListAttendanceUser() ([]*ListAttendance, error) {
	userId := s.user.ID
	const sqlGetByID = `SELECT ae.id, ev.description event_name, ae.is_delete, us.names user_name, ae.created_at, ae.updated_at
FROM [entity].[attendance_event_users]   ae
join [cfg].[events] ev on ev.id = ae.id_event 
join auth.users us on us.id = ae.id_user`
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
