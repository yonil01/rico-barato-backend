package comments

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

func newCommentSqlServerRepository(db *sqlx.DB, user *models.User, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *sqlserver) create(m *Comment) error {
	var id int
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	const sqlInsert = `INSERT INTO doc.comments (user_id, entity_id, value, start, is_delete, created_at, updated_at) VALUES (@user_id, @entity_id, @value, @start, @is_delete, @created_at, @updated_at) SELECT ID = convert(bigint, SCOPE_IDENTITY()) `
	stmt, err := s.DB.Prepare(sqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		sql.Named("user_id", m.UserId),
		sql.Named("entity_id", m.EntityId),
		sql.Named("value", m.Value),
		sql.Named("start", m.Start),
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
func (s *sqlserver) update(m *Comment) error {
	date := time.Now()
	m.UpdatedAt = date
	const sqlUpdate = `UPDATE doc.comments SET user_id = :user_id, entity_id = :entity_id, value = :value, start = :start, is_delete = :is_delete, updated_at = :updated_at WHERE id = :id `
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
	const sqlDelete = `DELETE FROM doc.comments WHERE id = :id `
	m := Comment{ID: id}
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
func (s *sqlserver) getByID(id int) (*Comment, error) {
	const sqlGetByID = `SELECT convert(nvarchar(50), id) id , user_id, entity_id, value, start, is_delete, created_at, updated_at FROM doc.comments  WITH (NOLOCK)  WHERE id = @id `
	mdl := Comment{}
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
func (s *sqlserver) getAll() ([]*Comment, error) {
	var ms []*Comment
	const sqlGetAll = `SELECT convert(nvarchar(50), id) id , user_id, entity_id, value, start, is_delete, created_at, updated_at FROM doc.comments  WITH (NOLOCK) `

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *sqlserver) getCommentByEntityID(id int) ([]*Comment, error) {
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
