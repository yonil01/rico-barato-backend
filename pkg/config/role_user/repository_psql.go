package role_user

import (
	"backend-ccff/internal/models"
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

func newRoleUserPsqlRepository(db *sqlx.DB, user *models.User, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *psql) create(m *RoleUser) error {
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	const psqlInsert = `INSERT INTO cfg.RoleUser (id ,name, description, event_date, is_disable, is_delete, user_id, created_at, updated_at) VALUES (:id ,:name, :description, :event_date, :is_disable, :is_delete, :user_id,:created_at, :updated_at) `
	rs, err := s.DB.NamedExec(psqlInsert, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("Dev-cff:108")
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *psql) update(m *RoleUser) error {
	date := time.Now()
	m.UpdatedAt = date
	const psqlUpdate = `UPDATE cfg.RoleUser SET name = :name, description = :description, event_date = :event_date, is_disable = :is_disable, is_delete = :is_delete, user_id = :user_id, updated_at = :updated_at WHERE id = :id `
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
func (s *psql) delete(id string) error {
	const psqlDelete = `DELETE FROM [auth].[users_roles] WHERE id = :id `
	m := RoleUser{ID: id}
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
func (s *psql) getByID(id string) (*RoleUser, error) {
	const psqlGetByID = `SELECT id , name, description, event_date, is_disable, is_delete, user_id, created_at, updated_at FROM cfg.RoleUser WHERE id = $1 `
	mdl := RoleUser{}
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
func (s *psql) getAll() ([]*RoleUser, error) {
	var ms []*RoleUser
	const psqlGetAll = ` SELECT id , name, description, event_date, is_disable, is_delete, user_id, created_at, updated_at FROM cfg.RoleUser `

	err := s.DB.Select(&ms, psqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *psql) getAllByUser(id string) ([]*RoleUser, error) {
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
