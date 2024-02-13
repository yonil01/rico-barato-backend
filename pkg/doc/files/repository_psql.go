package files

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

func newFilesPsqlRepository(db *sqlx.DB, user *models.User, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *psql) create(m *Files) error {

	const psqlInsert = `INSERT INTO doc.files (entity_id, path, type_document, type_entity, user_id, is_delete) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at, updated_at`
	stmt, err := s.DB.Prepare(psqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		m.EntityId,
		m.Path,
		m.TypeDocument,
		m.TypeEntity,
		m.UserId,
		m.IsDelete,
	).Scan(&m.ID, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		return err
	}
	return nil

}

// Update actualiza un registro en la BD
func (s *psql) update(m *Files) error {
	date := time.Now()
	m.UpdatedAt = date
	const psqlUpdate = `UPDATE doc.files SET entity_id = :entity_id, path = :path, type_document = :type_document, type_entity = :type_entity, user_id = :user_id, is_delete = :is_delete, updated_at = :updated_at WHERE id = :id `
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
	const psqlDelete = `DELETE FROM doc.files WHERE id = :id `
	m := Files{ID: id}
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
func (s *psql) getByID(id int) (*Files, error) {
	const psqlGetByID = `SELECT id , entity_id, path, type_document, type_entity, user_id, is_delete, created_at, updated_at FROM doc.files WHERE id = $1 `
	mdl := Files{}
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
func (s *psql) getAll() ([]*Files, error) {
	var ms []*Files
	const psqlGetAll = ` SELECT id , entity_id, path, type_document, type_entity, user_id, is_delete, created_at, updated_at FROM doc.files `

	err := s.DB.Select(&ms, psqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *psql) getFilesByEntityId(entityId int, typeEntity int) ([]*Files, error) {
	const psqlGetByID = `SELECT id , entity_id, path, type_document, type_entity, user_id, is_delete, created_at, updated_at FROM doc.files WHERE entity_id = $1 and type_entity = $2 `
	var ms []*Files
	err := s.DB.Select(&ms, psqlGetByID, entityId, typeEntity)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
