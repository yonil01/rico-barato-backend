package upload_metadata

import "gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/pkg/indra/upload_metadata"

type RequestProcess struct {
	Metadata []upload_metadata.Metadata `json:"metadata"`
}

type ResponseUploadMetadata struct {
	Error bool   `json:"error"`
	Data  string `json:"data"`
	Code  int    `json:"code"`
	Type  string `json:"type"`
	Msg   string `json:"msg"`
}
