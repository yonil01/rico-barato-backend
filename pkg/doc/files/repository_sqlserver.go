package files

import (
	"database/sql"
	"fmt"
	"time"

	"backend-comee/internal/models"
	"github.com/jmoiron/sqlx"
)

// sqlServer estructura de conexión a la BD de mssql
type sqlserver struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func newFilesSqlServerRepository(db *sqlx.DB, user *models.User, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *sqlserver) create(m *Files) error {
	var id int
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	const sqlInsert = `INSERT INTO doc.files (entity_id, path, type_document, type_entity, user_id, is_delete, created_at, updated_at) VALUES (@entity_id, @path, @type_document, @type_entity, @user_id, @is_delete, @created_at, @updated_at) SELECT ID = convert(bigint, SCOPE_IDENTITY()) `
	stmt, err := s.DB.Prepare(sqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		sql.Named("entity_id", m.EntityId),
		sql.Named("path", m.Path),
		sql.Named("type_document", m.TypeDocument),
		sql.Named("type_entity", m.TypeEntity),
		sql.Named("user_id", m.UserId),
		sql.Named("is_delete", m.IsDelete),
		sql.Named("created_at", m.CreatedAt),
		sql.Named("updated_at", m.UpdatedAt),
	).Scan(&id)
	if err != nil {
		return err
	}
	m.ID = int(id)
	return nil
}

// Update actualiza un registro en la BD
func (s *sqlserver) update(m *Files) error {
	date := time.Now()
	m.UpdatedAt = date
	const sqlUpdate = `UPDATE doc.files SET entity_id = :entity_id, path = :path, type_document = :type_document, type_entity = :type_entity, user_id = :user_id, is_delete = :is_delete, updated_at = :updated_at WHERE id = :id `
	rs, err := s.DB.NamedExec(sqlUpdate, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// Delete elimina un registro de la BD
func (s *sqlserver) delete(id int) error {
	const sqlDelete = `DELETE FROM doc.files WHERE id = :id `
	m := Files{ID: id}
	rs, err := s.DB.NamedExec(sqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// GetByID consulta un registro por su ID
func (s *sqlserver) getByID(id int) (*Files, error) {
	const sqlGetByID = `SELECT convert(nvarchar(50), id) id , entity_id, path, type_document, type_entity, user_id, is_delete, created_at, updated_at FROM doc.files  WITH (NOLOCK)  WHERE id = @id `
	mdl := Files{}
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
func (s *sqlserver) getAll() ([]*Files, error) {
	var ms []*Files
	const sqlGetAll = `SELECT convert(nvarchar(50), id) id , entity_id, path, type_document, type_entity, user_id, is_delete, created_at, updated_at FROM doc.files  WITH (NOLOCK) `

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *sqlserver) getFilesByEntityId(entityId int, typeEntity int) ([]*Files, error) {
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
