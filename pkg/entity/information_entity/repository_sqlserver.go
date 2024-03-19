package information_entity

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

func newInformationEntitySqlServerRepository(db *sqlx.DB, user *models.User, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *sqlserver) create(m *models.Entity) error {
	var id int
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	const sqlInsert = `INSERT INTO entity.information_entity (user_entity_id, name, description, telephone, mobile, location_x, location_y, is_block, is_delete, user_id, created_at, updated_at) VALUES (@user_entity_id, @name, @description, @telephone, @mobile, @location_x, @location_y, @is_block, @is_delete, @user_id, @created_at, @updated_at) SELECT ID = convert(bigint, SCOPE_IDENTITY()) `
	stmt, err := s.DB.Prepare(sqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		sql.Named("user_entity_id", m.UserEntityId),
		sql.Named("name", m.Name),
		sql.Named("description", m.Description),
		sql.Named("telephone", m.Telephone),
		sql.Named("mobile", m.Mobile),
		sql.Named("location_x", m.LocationX),
		sql.Named("location_y", m.LocationY),
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
func (s *sqlserver) update(m *models.Entity) error {
	date := time.Now()
	m.UpdatedAt = date
	const sqlUpdate = `UPDATE entity.information_entity SET user_entity_id = :user_entity_id, name = :name, description = :description, telephone = :telephone, mobile = :mobile, location_x = :location_x, location_y = :location_y, is_block = :is_block, is_delete = :is_delete, user_id = :user_id, updated_at = :updated_at WHERE id = :id `
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
	const sqlDelete = `DELETE FROM entity.information_entity WHERE id = :id `
	m := models.Entity{ID: id}
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
func (s *sqlserver) getByID(id int) (*models.Entity, error) {
	const sqlGetByID = `SELECT convert(nvarchar(50), id) id , user_entity_id, name, description, telephone, mobile, location_x, location_y, is_block, is_delete, user_id, created_at, updated_at FROM entity.information_entity  WITH (NOLOCK)  WHERE id = @id `
	mdl := models.Entity{}
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
func (s *sqlserver) getAll() ([]*models.Entity, error) {
	var ms []*models.Entity
	const sqlGetAll = `SELECT convert(nvarchar(50), id) id , user_entity_id, name, description, telephone, mobile, location_x, location_y, is_block, is_delete, user_id, created_at, updated_at FROM entity.information_entity  WITH (NOLOCK) `

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *sqlserver) getEntityByCoordinate(locationX string, locationY string, amount int) ([]*models.Entity, error) {
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

func (s *sqlserver) getByUserId(id string) ([]*models.Entity, error) {
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
