package indra

import (
	"github.com/jmoiron/sqlx"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/models"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/pkg/indra/upload_metadata"
)

type ServerIndra struct {
	SrvUploadMetadata upload_metadata.PortsServerUploadMetadata
}

func NewServerIndra(db *sqlx.DB, user *models.User, txID string) *ServerIndra {
	repoUploadMetadata := upload_metadata.FactoryStorage(db, user, txID)
	return &ServerIndra{
		SrvUploadMetadata: upload_metadata.NewUploadMetadataService(repoUploadMetadata, user, txID),
	}
}
