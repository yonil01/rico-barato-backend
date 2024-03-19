package information_entity

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

func newInformationEntityOrclRepository(db *sqlx.DB, user *models.User, txID string) *orcl {
	return &orcl{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *orcl) create(m *models.Entity) error {
	const osqlInsert = `INSERT INTO entity.information_entity (user_entity_id, name, description, telephone, mobile, location_x, location_y, is_block, is_delete, user_id, created_at, updated_at)  VALUES (:user_entity_id, :name, :description, :telephone, :mobile, :location_x, :location_y, :is_block, :is_delete, :user_id, sysdate, sysdate) RETURNING id into id, created_at into created_at, updated_at into updated_at `
	stmt, err := s.DB.Prepare(osqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		m.UserEntityId,
		m.Name,
		m.Description,
		m.Telephone,
		m.Mobile,
		m.LocationX,
		m.LocationY,
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
func (s *orcl) update(m *models.Entity) error {
	date := time.Now()
	m.UpdatedAt = date
	const osqlUpdate = `UPDATE entity.information_entity SET user_entity_id = :user_entity_id, name = :name, description = :description, telephone = :telephone, mobile = :mobile, location_x = :location_x, location_y = :location_y, is_block = :is_block, is_delete = :is_delete, user_id = :user_id, updated_at = :updated_at WHERE id = :id  `
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
	const osqlDelete = `DELETE FROM entity.information_entity WHERE id = :id `
	m := models.Entity{ID: id}
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
func (s *orcl) getByID(id int) (*models.Entity, error) {
	const osqlGetByID = `SELECT id , user_entity_id, name, description, telephone, mobile, location_x, location_y, is_block, is_delete, user_id, created_at, updated_at FROM entity.information_entity WHERE id = :1 `
	mdl := models.Entity{}
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
func (s *orcl) getAll() ([]*models.Entity, error) {
	var ms []*models.Entity
	const osqlGetAll = ` SELECT id , user_entity_id, name, description, telephone, mobile, location_x, location_y, is_block, is_delete, user_id, created_at, updated_at FROM entity.information_entity `

	err := s.DB.Select(&ms, osqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *orcl) getEntityByCoordinate(locationX string, locationY string, amount int) ([]*models.Entity, error) {
	const psqlGetByCoordinate = `
    SELECT id, user_entity_id, name, description, telephone, mobile, location_x, location_y, 
           is_block, is_delete, user_id, created_at, updated_at
    FROM entity.information_entity
    WHERE NOT is_block AND NOT is_delete
    ORDER BY (location_x::float - $1::float)^2 + (location_y::float - $2::float)^2
    LIMIT $3;`

	var mdl []*models.Entity

	err := s.DB.Select(&mdl, psqlGetByCoordinate, locationX, locationY, amount)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return mdl, nil
}

func (s *orcl) getByUserId(id string) ([]*models.Entity, error) {
	const psqlGetByID = `SELECT id , user_entity_id, name, description, telephone, mobile, location_x, location_y, is_block, is_delete, user_id, created_at, updated_at FROM entity.information_entity WHERE user_entity_id = $1 `
	mdl := []*models.Entity{}
	err := s.DB.Get(&mdl, psqlGetByID, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return mdl, err
	}
	return mdl, nil
}
