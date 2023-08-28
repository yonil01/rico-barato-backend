package roles

import (
	"backend-ccff/internal/helper"
	"backend-ccff/internal/logger"
	"backend-ccff/internal/models"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// psql estructura de conexión a la BD de postgresql
type psql struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func NewRolePsqlRepository(db *sqlx.DB, user *models.User, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *psql) Create(m *Role) error {
	m.IdUser = s.user.ID
	const sqlInsert = `INSERT INTO auth.roles (id ,name, description, sessions_allowed, see_all_users, id_user, created_at, updated_at) VALUES (:id ,:name, :description, :sessions_allowed, :see_all_users, :id_user, Now(), Now()) `
	_, err := s.DB.NamedExec(sqlInsert, &m)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't insert Role: %v", err)
		return err
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *psql) Update(m *Role) error {
	m.IdUser = s.user.ID
	const sqlUpdate = `UPDATE auth.roles SET name = :name, description = :description, sessions_allowed = :sessions_allowed, see_all_users = :see_all_users, id_user =:id_user, updated_at = Now() WHERE id = :id `
	rs, err := s.DB.NamedExec(sqlUpdate, &m)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't update Role: %v", err)
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("Dev-cff:108")
	}
	return nil
}

// Delete elimina un registro de la BD
func (s *psql) Delete(id string) error {
	m := Role{ID: id, IdUser: s.user.ID}
	const sqlUpdate = `UPDATE auth.roles SET id_user = :id_user, is_delete = true, updated_at = NOW(), deleted_at = NOW() WHERE id = :id `
	rs, err := s.DB.NamedExec(sqlUpdate, &m)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't delete Role: %v", err)
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("Dev-cff:108")
	}
	return nil
}

// GetByID consulta un registro por su ID
func (s *psql) GetByID(id string) (*Role, error) {
	const sqlGetByID = `SELECT id , name, description, sessions_allowed, see_all_users, created_at, updated_at FROM auth.roles  WHERE id = $1 AND is_delete = false `
	mdl := Role{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error.Printf(s.TxID, " - couldn't execute GetByID Role: %v", err)
		return &mdl, err
	}
	return &mdl, nil
}

// GetAll consulta todos los registros de la BD
func (s *psql) GetAll() ([]*Role, error) {
	var ms []*Role
	const sqlGetAll = `SELECT id , name, description, sessions_allowed, see_all_users, created_at, updated_at FROM auth.roles WHERE is_delete = false`

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error.Printf(s.TxID, " - couldn't execute GetAll auth.roles: %v", err)
		return ms, err
	}
	return ms, nil
}

// GetByUserID consulta un registro por su ID
func (s *psql) GetByUserID(userID string) ([]*Role, error) {
	var ms []*Role
	const sqlGetByID = `SELECT r.id , r.name, r.description, r.sessions_allowed, r.see_all_users, r.created_at, r.updated_at FROM auth.roles r 
			    		JOIN auth.users_roles ur ON r.id = ur.role_id WHERE ur.user_id = $1 AND r.is_delete = false AND ur.is_delete = false`
	err := s.DB.Select(&ms, sqlGetByID, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error.Printf(s.TxID, " - couldn't execute GetByUserID auth.roles: %v", err)
		return ms, err
	}
	return ms, nil
}

func (s *psql) GetRolesByProcessIDs(ProcessIDs []string) ([]*Role, error) {
	var ms []*Role
	const sqlGetRolesByProcessIDs = `SELECT a.id, a.name, a.description, a.sessions_allowed, b.process_id, a.see_all_users, a.created_at, a.updated_at FROM auth.roles a JOIN cfg.process_roles b ON  a.id = b.role_id WHERE b.process_id in (%s) AND a.is_delete = false AND b.is_delete = false`

	err := s.DB.Select(&ms, fmt.Sprintf(sqlGetRolesByProcessIDs, helper.SliceToString(ProcessIDs)))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error.Printf(s.TxID, " - couldn't execute GetRolesByProcessIDs auth.roles: %v", err)
		return ms, err
	}
	return ms, nil
}

func (s *psql) GetRolesByIDs(IDs []string) ([]*Role, error) {
	var ms []*Role
	const sqlGetRolesByIDs = `SELECT a.id, a.name, a.description, a.sessions_allowed, a.see_all_users, a.created_at, a.updated_at FROM auth.roles a WHERE a.id in (%s) AND a.is_delete = false `

	err := s.DB.Select(&ms, fmt.Sprintf(sqlGetRolesByIDs, helper.SliceToString(IDs)))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error.Printf(s.TxID, " - couldn't execute GetRolesByIDs auth.roles: %v", err)
		return ms, err
	}
	return ms, nil
}

func (s *psql) GetByUserIDs(userIDs []string) ([]*Role, error) {
	var ms []*Role
	const sqlGetByID = `SELECT r.id, r.name, r.description, r.sessions_allowed, ur.user_id, r.see_all_users, r.created_at, r.updated_at  FROM auth.roles r JOIN auth.users_roles ur ON r.id = ur.role_id WHERE ur.user_id IN (%s) AND r.is_delete = false AND ur.is_delete = false `
	query := fmt.Sprintf(sqlGetByID, helper.SliceToString(userIDs))
	err := s.DB.Select(&ms, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error.Printf(s.TxID, " - couldn't execute GetByUserIDs auth.roles: %v", err)
		return ms, err
	}
	return ms, nil
}

// GetRolesByProjectID consulta todos los registros de la BD
func (s *psql) GetRolesByProjectID(projectID string) ([]*Role, error) {
	var ms []*Role
	const sqlGetRolesByprojectID = `SELECT r.id , r.name, r.description, r.sessions_allowed, r.see_all_users, r.created_at, r.updated_at 
	FROM auth.roles r JOIN auth.roles_projects rp ON (r.id = rp.role_id)
	WHERE rp.project = '%s' AND r.is_delete = false AND rp.is_delete = false`

	err := s.DB.Select(&ms, fmt.Sprintf(sqlGetRolesByprojectID, projectID))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error.Printf(s.TxID, " - couldn't execute GetRolesByProjectID auth.roles: %v", err)
		return ms, err
	}
	return ms, nil
}

func (s *psql) GetRolesByQueueId(queueId string) ([]*Role, error) {
	var ms []*Role
	const sqlGetRolesByQueueId = `SELECT r.id, r.name, r.description, r.sessions_allowed, r.id_user as user_id, r.see_all_users, r.created_at, r.updated_at
	FROM cfg.queues_roles qr JOIN auth.roles r ON r.id = qr.role_id WHERE qr.queue_id = $1`

	err := s.DB.Select(&ms, sqlGetRolesByQueueId, queueId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error.Printf(s.TxID, " - couldn't execute GetRolesByQueueId auth.roles: %v", err)
		return ms, err
	}
	return ms, nil
}

func (s *psql) GetRoleByName(name string) (*Role, error) {
	var ms *Role
	const sqlGetRolesByQueueId = `select convert(nvarchar(50), id) id , name, description, sessions_allowed, see_all_users, created_at, updated_at FROM auth.roles  WITH (NOLOCK) WHERE  UPPER(description)  = Upper($1)`

	err := s.DB.Get(&ms, sqlGetRolesByQueueId, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error.Printf(s.TxID, " - couldn't execute GetRolesByQueueId auth.roles: %v", err)
		return ms, err
	}
	return ms, nil
}
