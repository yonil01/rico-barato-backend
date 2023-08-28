package modules

import (
	"backend-ccff/internal/helper"
	"backend-ccff/internal/logger"
	"backend-ccff/internal/models"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

// psql estructura de conexión a la BD de postgresql
type psql struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func NewModulePsqlRepository(db *sqlx.DB, user *models.User, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *psql) Create(m *Module) error {
	const sqlInsert = `INSERT INTO auth.modules (id ,name, description, class, id_user, created_at, updated_at) VALUES (:id ,:name, :description, :class,:id_user ,Now(), Now()) `
	m.IdUser = s.user.ID
	_, err := s.DB.NamedExec(sqlInsert, &m)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't insert Module: %v", err)
		return err
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *psql) Update(m *Module) error {
	const sqlUpdate = `UPDATE auth.modules SET name = :name, description = :description, class = :class, id_user =:id_user, updated_at = Now() WHERE id = :id `
	m.IdUser = s.user.ID
	rs, err := s.DB.NamedExec(sqlUpdate, &m)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't update Module: %v", err)
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("Dev-cff:108")
	}
	return nil
}

// Delete elimina un registro de la BD
func (s *psql) Delete(id string) error {
	m := Module{ID: id, IdUser: s.user.ID}
	const sqlDelete = `UPDATE auth.modules SET is_delete = true, id_user =:id_user, updated_at = GetDate(), deleted_at = GetDate() WHERE id = :id `
	rs, err := s.DB.NamedExec(sqlDelete, &m)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't delete Module: %v", err)
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("Dev-cff:108")
	}
	return nil
}

// GetByID consulta un registro por su ID
func (s *psql) GetByID(id string) (*Module, error) {
	const sqlGetByID = `SELECT id , name, description, class, created_at, updated_at FROM auth.modules WHERE id = $1 AND is_delete = false`
	mdl := Module{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error.Printf(s.TxID, " - couldn't execute GetByID Module: %v", err)
		return &mdl, err
	}
	return &mdl, nil
}

// GetAll consulta todos los registros de la BD
func (s *psql) GetAll() ([]*Module, error) {
	var ms []*Module
	const sqlGetAll = `select el.id, mm.name, mm.description, mm.class, mm.created_at, mm.updated_at from [auth].[modules_components_elements] el
				join [auth].[modules_components] mc on mc.id = el.component_id
				join auth.modules mm on mm.id = mc.module_id WHERE mm.is_delete = false`
	query := sqlGetAll
	err := s.DB.Select(&ms, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error.Printf(s.TxID, " - couldn't execute GetAll auth.modules: %v", err)
		return ms, err
	}
	return ms, nil
}

// GetModulesByRoles consulta todos los registros de la BD
func (s *psql) GetModulesByRoles(roleIDs []string, ids []string, typeArg int) ([]*Module, error) {
	var ms []*Module
	const sqlGetModulesByRoles = `SELECT DISTINCT m.id, m.name, m.description, m.class, m.created_at, m.updated_at FROM auth.modules m  
				JOIN auth.modules_components c  ON (m.id = c.module_id)
				JOIN auth.modules_components_elements e ON (c.id = e.component_id)
				JOIN auth.roles_elements re  ON (e.id = re.element_id)
				WHERE re.role_id in (%s) AND m.type = $1 
				AND m.is_delete = false AND c.is_delete = false AND e.is_delete = false AND re.is_delete = false
`
	// TODO ADD IDS
	query := fmt.Sprintf(sqlGetModulesByRoles, helper.SliceToString(roleIDs))
	err := s.DB.Select(&ms, query, typeArg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error.Printf(s.TxID, " - couldn't execute GetModulesByRoles auth.modules: %v", err)
		return ms, err
	}
	return ms, nil
}

func (s *psql) GetModulesRole(roleId string) ([]*ModuleRole, error) {
	const sqlGetByID = `select id, role_id, element_id, id_user, created_at, updated_at  from [auth].[roles_elements] where role_id = @role_id`
	var mdl []*ModuleRole
	err := s.DB.Select(&mdl, sqlGetByID, sql.Named("role_id", roleId))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error.Printf(s.TxID, " - couldn't execute GetByID Module: %v", err)
		return mdl, err
	}
	return mdl, nil
}

func (s *psql) DeleteModuleUser(id string) error {
	m := Module{ID: id, IdUser: s.user.ID}
	const sqlDelete = `DELETE from [auth].[roles_elements] WHERE id = :id `
	rs, err := s.DB.NamedExec(sqlDelete, &m)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't delete Module: %v", err)
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("Dev-cff:108")
	}
	return nil
}

func (s *psql) CreateModuleRole(m *ModuleRole) error {
	const sqlInsert = `INSERT INTO auth.roles_elements (id ,role_id, element_id, id_user, created_at, updated_at) VALUES (:id ,:role_id, :element_id, :id_user, GetDate(), GetDate()) `
	m.IdUser = s.user.ID
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	_, err := s.DB.NamedExec(sqlInsert, &m)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't insert Module: %v", err)
		return err
	}
	return nil
}
