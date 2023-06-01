package upload_metadata

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/logger"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/models"
)

type psql struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func newUploadMetadataPsqlRepository(db *sqlx.DB, user *models.User, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

func (s *psql) updateMetadata(metadata []Metadata) (int, error) {
	data, err := json.Marshal(metadata)
	json := string(data)

	const psqlExecuteSP = `select * from public.pqr_create_autofills_informacion_personal_array($1)`

	m := ValueData{}
	err1 := s.DB.Get(&m, psqlExecuteSP, json)
	if err != nil {
		logger.Error.Printf("preparando la sentencia ExecuteSP: %v", err)
		return 0, err1
	}
	return 1, nil
}

type ValueData struct {
	Value string `json:"value" db:"value" valid:"required"`
}
