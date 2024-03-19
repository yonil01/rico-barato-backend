package user_entity

import (
	"database/sql"
	"fmt"
	"time"

	"backend-comee/internal/models"
	"github.com/jmoiron/sqlx"
)

// psql estructura de conexión a la BD de postgresql
type psql struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func newUserEntityPsqlRepository(db *sqlx.DB, user *models.User, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *psql) create(m *UserEntity) error {
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	const psqlInsert = `INSERT INTO auth.user_entity (id ,dni, name, lastname, email, password, is_block, is_delete, user_id, created_at, updated_at) VALUES (:id ,:dni, :name, :lastname, :email, :password, :is_block, :is_delete, :user_id,:created_at, :updated_at) `
	rs, err := s.DB.NamedExec(psqlInsert, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *psql) update(m *UserEntity) error {
	date := time.Now()
	m.UpdatedAt = date
	const psqlUpdate = `UPDATE auth.user_entity SET dni = :dni, name = :name, lastname = :lastname, email = :email, password = :password, is_block = :is_block, is_delete = :is_delete, user_id = :user_id, updated_at = :updated_at WHERE id = :id `
	rs, err := s.DB.NamedExec(psqlUpdate, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// Delete elimina un registro de la BD
func (s *psql) delete(id string) error {
	const psqlDelete = `DELETE FROM auth.user_entity WHERE id = :id `
	m := UserEntity{ID: id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// GetByID consulta un registro por su ID
func (s *psql) getByID(id string) (*UserEntity, error) {
	const psqlGetByID = `SELECT id , dni, name, lastname, email, password, is_block, is_delete, user_id, created_at, updated_at FROM auth.user_entity WHERE id = $1 `
	mdl := UserEntity{}
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
func (s *psql) getAll() ([]*UserEntity, error) {
	var ms []*UserEntity
	const psqlGetAll = ` SELECT id , dni, name, lastname, email, password, is_block, is_delete, user_id, created_at, updated_at FROM auth.user_entity `

	err := s.DB.Select(&ms, psqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *psql) login(email string, password string) (*UserEntity, error) {
	const psqlGetByID = `SELECT id , dni, name, lastname, email, password, is_block, is_delete, user_id, created_at, updated_at FROM auth.user_entity WHERE email = $1`
	mdl := UserEntity{}
	err := s.DB.Get(&mdl, psqlGetByID, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}
