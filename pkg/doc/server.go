package doc

import (
	"backend-comee/internal/models"
	"backend-comee/pkg/doc/comments"
	"backend-comee/pkg/doc/files"
	"github.com/jmoiron/sqlx"
)

type ServerEntity struct {
	Files   files.PortsServerFiles
	Comment comments.PortsServerComment
}

func NewServerEntity(db *sqlx.DB, user *models.User, txID string) *ServerEntity {
	repoFiles := files.FactoryStorage(db, user, txID)
	repoComments := comments.FactoryStorage(db, user, txID)
	return &ServerEntity{
		Files:   files.NewFilesService(repoFiles, user, txID),
		Comment: comments.NewCommentService(repoComments, user, txID),
	}
}
