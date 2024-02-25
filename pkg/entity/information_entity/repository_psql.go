package information_entity

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

func newInformationEntityPsqlRepository(db *sqlx.DB, user *models.User, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *psql) create(m *models.Entity) error {

	const psqlInsert = `INSERT INTO entity.information_entity (user_entity_id, name, description, telephone, mobile, location_x, location_y, is_block, is_delete, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id, created_at, updated_at`
	stmt, err := s.DB.Prepare(psqlInsert)
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
func (s *psql) update(m *models.Entity) error {
	date := time.Now()
	m.UpdatedAt = date
	const psqlUpdate = `UPDATE entity.information_entity SET user_entity_id = :user_entity_id, name = :name, description = :description, telephone = :telephone, mobile = :mobile, location_x = :location_x, location_y = :location_y, is_block = :is_block, is_delete = :is_delete, user_id = :user_id, updated_at = :updated_at WHERE id = :id `
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
	const psqlDelete = `DELETE FROM entity.information_entity WHERE id = :id `
	m := models.Entity{ID: id}
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
func (s *psql) getByID(id int) (*models.Entity, error) {
	const psqlGetByID = `SELECT id , user_entity_id, name, description, telephone, mobile, location_x, location_y, is_block, is_delete, user_id, created_at, updated_at FROM entity.information_entity WHERE id = $1 `
	mdl := models.Entity{}
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
func (s *psql) getAll() ([]*models.Entity, error) {
	var ms []*models.Entity
	const psqlGetAll = ` SELECT id , user_entity_id, name, description, telephone, mobile, location_x, location_y, is_block, is_delete, user_id, created_at, updated_at FROM entity.information_entity `

	err := s.DB.Select(&ms, psqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *psql) getEntityByCoordinate(locationX string, locationY string, amount int) ([]*models.Entity, error) {
	const psqlGetByCoordinate = `
    SELECT id, user_entity_id, name, description, telephone, mobile, location_x, location_y, 
           is_block, is_delete, user_id, created_at, updated_at
    FROM entity.information_entity
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
