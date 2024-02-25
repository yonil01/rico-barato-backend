package comments

import (
	"database/sql"
	"fmt"
	"time"

	"backend-comee/internal/models"
	"github.com/jmoiron/sqlx"
)

// Orcl estructura de conexión a la BD de Oracle
type orcl struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func newCommentOrclRepository(db *sqlx.DB, user *models.User, txID string) *orcl {
	return &orcl{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *orcl) create(m *Comment) error {
	const osqlInsert = `INSERT INTO doc.comments (user_id, entity_id, value, start, is_delete, created_at, updated_at)  VALUES (:user_id, :entity_id, :value, :start, :is_delete, sysdate, sysdate) RETURNING id into id, created_at into created_at, updated_at into updated_at `
	stmt, err := s.DB.Prepare(osqlInsert)
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
func (s *orcl) update(m *Comment) error {
	date := time.Now()
	m.UpdatedAt = date
	const osqlUpdate = `UPDATE doc.comments SET user_id = :user_id, entity_id = :entity_id, value = :value, start = :start, is_delete = :is_delete, updated_at = :updated_at WHERE id = :id  `
	rs, err := s.DB.NamedExec(osqlUpdate, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// Delete elimina un registro de la BD
func (s *orcl) delete(id int) error {
	const osqlDelete = `DELETE FROM doc.comments WHERE id = :id `
	m := Comment{ID: id}
	rs, err := s.DB.NamedExec(osqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// GetByID consulta un registro por su ID
func (s *orcl) getByID(id int) (*Comment, error) {
	const osqlGetByID = `SELECT id , user_id, entity_id, value, start, is_delete, created_at, updated_at FROM doc.comments WHERE id = :1 `
	mdl := Comment{}
	err := s.DB.Get(&mdl, osqlGetByID, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

// GetAll consulta todos los registros de la BD
func (s *orcl) getAll() ([]*Comment, error) {
	var ms []*Comment
	const osqlGetAll = ` SELECT id , user_id, entity_id, value, start, is_delete, created_at, updated_at FROM doc.comments `

	err := s.DB.Select(&ms, osqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *orcl) getCommentByEntityID(id int) ([]*Comment, error) {
	const psqlGetByID = `SELECT id , user_id, entity_id, value, start, is_delete, created_at, updated_at FROM doc.comments WHERE entity_id = $1 `
	var mdl []*Comment
	err := s.DB.Get(&mdl, psqlGetByID, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return mdl, err
	}
	return mdl, nil
}
