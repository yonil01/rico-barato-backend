package api

import "gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/dbx"

func Start(port int, app string, loggerHttp bool, allowedOrigins string) {

	db := dbx.GetConnection()
	defer db.Close()

	r := routes(db, loggerHttp, allowedOrigins)
	server := newServer(port, app, r)
	server.Start()
}
