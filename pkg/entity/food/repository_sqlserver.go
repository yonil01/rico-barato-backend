package food

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

func newFoodSqlServerRepository(db *sqlx.DB, user *models.User, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *sqlserver) create(m *models.Food) error {
	var id int
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	const sqlInsert = `INSERT INTO entity.food (entity_id, name, description, price, status, is_block, is_delete, user_id, created_at, updated_at) VALUES (@entity_id, @name, @description, @price, @status, @is_block, @is_delete, @user_id, @created_at, @updated_at) SELECT ID = convert(bigint, SCOPE_IDENTITY()) `
	stmt, err := s.DB.Prepare(sqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		sql.Named("entity_id", m.EntityId),
		sql.Named("name", m.Name),
		sql.Named("description", m.Description),
		sql.Named("price", m.Price),
		sql.Named("status", m.Status),
		sql.Named("is_block", m.IsBlock),
		sql.Named("is_delete", m.IsDelete),
		sql.Named("user_id", m.UserId),
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
func (s *sqlserver) update(m *models.Food) error {
	date := time.Now()
	m.UpdatedAt = date
	const sqlUpdate = `UPDATE entity.food SET entity_id = :entity_id, name = :name, description = :description, price = :price, status = :status, is_block = :is_block, is_delete = :is_delete, user_id = :user_id, updated_at = :updated_at WHERE id = :id `
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
	const sqlDelete = `DELETE FROM entity.food WHERE id = :id `
	m := models.Food{ID: id}
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
func (s *sqlserver) getByID(id int) (*models.Food, error) {
	const sqlGetByID = `SELECT convert(nvarchar(50), id) id , entity_id, name, description, price, status, is_block, is_delete, user_id, created_at, updated_at FROM entity.food  WITH (NOLOCK)  WHERE id = @id `
	mdl := models.Food{}
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
func (s *sqlserver) getAll() ([]*models.Food, error) {
	var ms []*models.Food
	const sqlGetAll = `SELECT convert(nvarchar(50), id) id , entity_id, name, description, price, status, is_block, is_delete, user_id, created_at, updated_at FROM entity.food  WITH (NOLOCK) `

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *sqlserver) getFoodsByEntityId(entityId int) ([]*models.Food, error) {
	const psqlGetByID = `SELECT id , entity_id, name, description, price, status, is_block, is_delete, user_id, created_at, updated_at FROM entity.food WHERE entity_id = $1 `
	var ms []*models.Food
	err := s.DB.Select(&ms, psqlGetByID, entityId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
