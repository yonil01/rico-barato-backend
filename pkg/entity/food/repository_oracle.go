package food

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

func newFoodOrclRepository(db *sqlx.DB, user *models.User, txID string) *orcl {
	return &orcl{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *orcl) create(m *Food) error {
	const osqlInsert = `INSERT INTO entity.food (entity_id, name, description, price, status, is_block, is_delete, user_id, created_at, updated_at)  VALUES (:entity_id, :name, :description, :price, :status, :is_block, :is_delete, :user_id, sysdate, sysdate) RETURNING id into id, created_at into created_at, updated_at into updated_at `
	stmt, err := s.DB.Prepare(osqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		m.EntityId,
		m.Name,
		m.Description,
		m.Price,
		m.Status,
		m.IsBlock,
		m.IsDelete,
		m.UserId,
	).Scan(&m.ID, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *orcl) update(m *Food) error {
	date := time.Now()
	m.UpdatedAt = date
	const osqlUpdate = `UPDATE entity.food SET entity_id = :entity_id, name = :name, description = :description, price = :price, status = :status, is_block = :is_block, is_delete = :is_delete, user_id = :user_id, updated_at = :updated_at WHERE id = :id  `
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
	const osqlDelete = `DELETE FROM entity.food WHERE id = :id `
	m := Food{ID: id}
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
func (s *orcl) getByID(id int) (*Food, error) {
	const osqlGetByID = `SELECT id , entity_id, name, description, price, status, is_block, is_delete, user_id, created_at, updated_at FROM entity.food WHERE id = :1 `
	mdl := Food{}
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
func (s *orcl) getAll() ([]*Food, error) {
	var ms []*Food
	const osqlGetAll = ` SELECT id , entity_id, name, description, price, status, is_block, is_delete, user_id, created_at, updated_at FROM entity.food `

	err := s.DB.Select(&ms, osqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *orcl) getFoodsByEntityId(entityId int) ([]*Food, error) {
	const psqlGetByID = `SELECT id , entity_id, name, description, price, status, is_block, is_delete, user_id, created_at, updated_at FROM entity.food WHERE entity_id = $1 `
	var ms []*Food
	err := s.DB.Select(&ms, psqlGetByID, entityId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
