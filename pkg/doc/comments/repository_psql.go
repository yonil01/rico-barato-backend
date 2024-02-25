package comments

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

func newCommentPsqlRepository(db *sqlx.DB, user *models.User, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *psql) create(m *Comment) error {

	const psqlInsert = `INSERT INTO doc.comments (user_id, entity_id, value, start, is_delete) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at`
	stmt, err := s.DB.Prepare(psqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		m.UserId,
		m.EntityId,
		m.Value,
		m.Start,
		m.IsDelete,
	).Scan(&m.ID, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		return err
	}
	return nil

}

// Update actualiza un registro en la BD
func (s *psql) update(m *Comment) error {
	date := time.Now()
	m.UpdatedAt = date
	const psqlUpdate = `UPDATE doc.comments SET user_id = :user_id, entity_id = :entity_id, value = :value, start = :start, is_delete = :is_delete, updated_at = :updated_at WHERE id = :id `
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
func (s *psql) delete(id int) error {
	const psqlDelete = `DELETE FROM doc.comments WHERE id = :id `
	m := Comment{ID: id}
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
func (s *psql) getByID(id int) (*Comment, error) {
	const psqlGetByID = `SELECT id , user_id, entity_id, value, start, is_delete, created_at, updated_at FROM doc.comments WHERE id = $1 `
	mdl := Comment{}
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
func (s *psql) getAll() ([]*Comment, error) {
	var ms []*Comment
	const psqlGetAll = ` SELECT id , user_id, entity_id, value, start, is_delete, created_at, updated_at FROM doc.comments `

	err := s.DB.Select(&ms, psqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *psql) getCommentByEntityID(id int) ([]*Comment, error) {
	const psqlGetByID = `SELECT id , user_id, entity_id, value, start, is_delete, created_at, updated_at FROM doc.comments WHERE entity_id = $1 `
	var mdl []*Comment
	err := s.DB.Select(&mdl, psqlGetByID, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return mdl, err
	}
	return mdl, nil
}
