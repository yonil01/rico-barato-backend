package reniec

import (
	"github.com/jmoiron/sqlx"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/models"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/pkg/indra/upload_metadata"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/pkg/reniec/dni"
)

type ServerReniec struct {
	Dni dni.PortsServerReniec
}

func NewServerReniec(db *sqlx.DB, user *models.User, txID string) *ServerReniec {
	repoDni := upload_metadata.FactoryStorage(db, user, txID)
	return &ServerReniec{
		Dni: dni.NewReniecService(repoDni, user, txID),
	}
}
