package files

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

func newFilesOrclRepository(db *sqlx.DB, user *models.User, txID string) *orcl {
	return &orcl{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *orcl) create(m *models.Files) error {
	const osqlInsert = `INSERT INTO doc.files (entity_id, path, type_document, type_entity, user_id, is_delete, created_at, updated_at)  VALUES (:entity_id, :path, :type_document, :type_entity, :user_id, :is_delete, sysdate, sysdate) RETURNING id into id, created_at into created_at, updated_at into updated_at `
	stmt, err := s.DB.Prepare(osqlInsert)
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
func (s *orcl) update(m *models.Files) error {
	date := time.Now()
	m.UpdatedAt = date
	const osqlUpdate = `UPDATE doc.files SET entity_id = :entity_id, path = :path, type_document = :type_document, type_entity = :type_entity, user_id = :user_id, is_delete = :is_delete, updated_at = :updated_at WHERE id = :id  `
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
	const osqlDelete = `DELETE FROM doc.files WHERE id = :id `
	m := models.Files{ID: id}
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
func (s *orcl) getByID(id int) (*models.Files, error) {
	const osqlGetByID = `SELECT id , entity_id, path, type_document, type_entity, user_id, is_delete, created_at, updated_at FROM doc.files WHERE id = :1 `
	mdl := models.Files{}
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
func (s *orcl) getAll() ([]*models.Files, error) {
	var ms []*models.Files
	const osqlGetAll = ` SELECT id , entity_id, path, type_document, type_entity, user_id, is_delete, created_at, updated_at FROM doc.files `

	err := s.DB.Select(&ms, osqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *orcl) getFilesByEntityId(entityId int, typeEntity int) ([]*models.Files, error) {
	const psqlGetByID = `SELECT id , entity_id, path, type_document, type_entity, user_id, is_delete, created_at, updated_at FROM doc.files WHERE entity_id = $1 and type_entity = $2 `
	var ms []*models.Files
	err := s.DB.Select(&ms, psqlGetByID, entityId, typeEntity)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
