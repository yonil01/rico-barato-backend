package role_user

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

func newRoleUserSqlServerRepository(db *sqlx.DB, user *models.User, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *sqlserver) create(m *RoleUser) error {
	m.IdUser = s.user.ID
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	const sqlInsert = `INSERT INTO auth.users_roles (id ,user_id, role_id, id_user, is_delete, created_at, updated_at) VALUES (:id , :user_id, :role_id, :id_user, 0, :created_at, :updated_at) `
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
func (s *sqlserver) update(m *RoleUser) error {
	date := time.Now()
	m.UpdatedAt = date
	const sqlUpdate = `UPDATE auth.users_roles SET name = :name, description = :description, event_date = :event_date, is_disable = :is_disable, is_delete = :is_delete, user_id = :user_id, updated_at = :updated_at WHERE id = :id `
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
	const sqlDelete = `DELETE FROM auth.users_roles WHERE id = :id `
	m := RoleUser{ID: id}
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
func (s *sqlserver) getByID(id string) (*RoleUser, error) {
	const sqlGetByID = `SELECT convert(nvarchar(50), id) id , name, description, event_date, is_disable, is_delete, user_id, created_at, updated_at FROM cfg.RoleUser  WITH (NOLOCK)  WHERE user_id = @user_id `
	mdl := RoleUser{}
	err := s.DB.Get(&mdl, sqlGetByID, sql.Named("user_id", id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

// GetAll consulta todos los registros de la BD
func (s *sqlserver) getAll() ([]*RoleUser, error) {
	var ms []*RoleUser
	const sqlGetAll = `SELECT convert(nvarchar(50), id) id , convert(nvarchar(50), user_id) user_id, convert(nvarchar(50), role_id) role_id, convert(nvarchar(50), id_user) id_user, created_at, updated_at FROM auth.users_roles  WITH (NOLOCK) `

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *sqlserver) getAllByUser(id string) ([]*RoleUser, error) {
	var ms []*RoleUser
	const sqlGetAll = `SELECT convert(nvarchar(50), id) id , convert(nvarchar(50), user_id) user_id, convert(nvarchar(50), role_id) role_id, convert(nvarchar(50), id_user) id_user, created_at, updated_at FROM auth.users_roles  WITH (NOLOCK) WHERE user_id = @user_id`
	err := s.DB.Select(&ms, sqlGetAll, sql.Named("user_id", id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
