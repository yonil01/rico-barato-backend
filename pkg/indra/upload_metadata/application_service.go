package upload_metadata

import "gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/models"

type PortsServerUploadMetadata interface {
	UpdateMetadata(metadata []Metadata) (int, error)
	GetIdsAutofillValueByEntityAttributeAndValue(typeInput string, inputData string) ([]*Metadata, error)
}

type service struct {
	repository ServicesUploadMetadataRepository
	user       *models.User
	txID       string
}

func NewUploadMetadataService(repository ServicesUploadMetadataRepository, user *models.User, txID string) PortsServerUploadMetadata {
	return &service{
		repository: repository,
		user:       user,
		txID:       txID,
	}
}

func (s *service) UpdateMetadata(metadata []Metadata) (int, error) {
	return s.repository.updateMetadata(metadata)
}

func (s *service) GetIdsAutofillValueByEntityAttributeAndValue(typeInput string, inputData string) ([]*Metadata, error) {
	return s.repository.GetIdsAutofillValueByEntityAttributeAndValue(typeInput, inputData)
}
