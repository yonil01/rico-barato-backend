package users

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

func newUsersSqlServerRepository(db *sqlx.DB, user *models.User, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *sqlserver) create(m *Users) error {
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	const sqlInsert = `INSERT INTO auth.users (id ,username, code_student, dni, names, lastname_father, lastname_mother, email, password, is_delete, is_block, created_at, updated_at) VALUES (:id ,:username, :code_student, :dni, :names, :lastname_father, :lastname_mother, :email, :password, :is_delete, :is_block, :created_at, :updated_at) `
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
func (s *sqlserver) update(m *Users) error {
	date := time.Now()
	m.UpdatedAt = date
	const sqlUpdate = `UPDATE auth.users SET username = :username, code_student = :code_student, dni = :dni, names = :names, lastname_father = :lastname_father, lastname_mother = :lastname_mother, email = :email, password = :password, is_delete = :is_delete, is_block = :is_block, updated_at = :updated_at WHERE id = :id `
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
	const sqlDelete = `DELETE FROM auth.users WHERE id = :id `
	m := Users{ID: id}
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
func (s *sqlserver) getByID(id string) (*Users, error) {
	const sqlGetByID = `SELECT convert(nvarchar(50), id) id , username, code_student, dni, names, lastname_father, lastname_mother, email, password, is_delete, is_block, created_at, updated_at FROM auth.users  WITH (NOLOCK)  WHERE id = @id `
	mdl := Users{}
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
func (s *sqlserver) getAll() ([]*Users, error) {
	var ms []*Users
	const sqlGetAll = `SELECT convert(nvarchar(50), id) id , username, code_student, dni, names, lastname_father, lastname_mother, email, password, is_delete, is_block, created_at, updated_at FROM auth.users  WITH (NOLOCK) `

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *sqlserver) getByCodeStudent(codeStudent string) (*Users, error) {
	const sqlGetByID = `SELECT convert(nvarchar(50), id) id , username, code_student, dni, names, lastname_father, lastname_mother, email, password, is_delete, is_block, created_at, updated_at FROM auth.users  WITH (NOLOCK)  WHERE code_student = @code_student `
	mdl := Users{}
	err := s.DB.Get(&mdl, sqlGetByID, sql.Named("code_student", codeStudent))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}
