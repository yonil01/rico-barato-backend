package report

import (
	"backend-ccff/internal/logger"
	"backend-ccff/internal/models"
	"bytes"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// sqlServer estructura de conexión a la BD de mssql
type sqlserver struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func (s *sqlserver) ExecuteSP(procedure string, parameters map[string]interface{}, option int) ([]map[string]interface{}, error) {
	const sqlserverExecuteSP = `execute `
	rs := make([]map[string]interface{}, 0)

	r := bytes.Buffer{}
	r.WriteString(fmt.Sprintf(`%s %s `, sqlserverExecuteSP, procedure))

	for i, v := range parameters {
		if len(v.(string)) > 0 {
			_, err := r.WriteString(fmt.Sprintf(`@%s = '%s',`, i, v))
			if err != nil {
				logger.Error.Printf("agregando parametros a la ejecucion del SP en sqlserverExecuteSP: %v", err)
				return rs, err
			}
		}

	}
	if option == 1 {
		_, err := r.WriteString(fmt.Sprintf(`@user_id = '%s',`, s.user.ID))
		if err != nil {
			logger.Error.Printf("agregando parametro usuario a la ejecucion del SP en sqlserverExecuteSP: %v", err)
			return rs, err
		}
	}
	r.Truncate(r.Len() - 1)
	stmt, err := s.DB.Prepare(r.String())
	if err != nil {

		logger.Error.Printf("preparando consulta sqlserverExecuteSP: %s, %t", r.String(), err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		logger.Error.Printf("ejecutando sqlserverExecuteSP user: %s, %t", r.String(), err)
		return rs, err
	}
	defer rows.Close()

	cols, _ := rows.Columns()
	for rows.Next() {
		r := make(map[string]interface{})
		// Create a slice of interface{}'s to represent each column,
		// and a second slice to contain pointers to each item in the columns slice.
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}
		// Scan the result into the column pointers...
		if err = rows.Scan(columnPointers...); err != nil {
			logger.Error.Printf("no se pudo escanear las columnas de la consulta sqlserverGetInfoktg: %t", err)
			return nil, err
		}

		// Create our map, and retrieve the value for each column from the pointers slice,
		// storing it in the map with the name of the column as the key.
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			r[colName] = *val
		}
		rs = append(rs, r)
	}
	return rs, nil
}

func (s *sqlserver) ExecuteSPBYDocumentID(procedure string, documentID int64) (int, error) {
	const sqlserverExecuteSP = `EXECUTE %s %s, %s`
	var res int
	sqlExecute := fmt.Sprintf(sqlserverExecuteSP, procedure, "@document", "@user")
	stmt, err := s.DB.Prepare(sqlExecute)
	if err != nil {
		logger.Error.Printf("preparando la sentencia ExecuteSP: %v", err)
		return 0, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		sql.Named("document", documentID),
		sql.Named("user", s.user.ID),
	).Scan(&res)
	if err != nil {
		logger.Error.Printf("***ejecutando la sentencia ExecuteSP: %v", err)
		return 0, err
	}
	return res, nil
}

func NewReportSqlServerRepository(db *sqlx.DB, user *models.User, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		user: user,
		TxID: txID,
	}
}
