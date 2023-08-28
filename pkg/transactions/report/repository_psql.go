package report

import (
	"backend-ccff/internal/logger"
	"backend-ccff/internal/models"
	"bytes"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// psql estructura de conexión a la BD de mssql
type psql struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func (s *psql) ExecuteSP(procedure string, parameters map[string]interface{}, option int) ([]map[string]interface{}, error) {
	const psqlExecuteSP = `select * from %s `
	rs := make([]map[string]interface{}, 0)

	r := bytes.Buffer{}
	r.WriteString(fmt.Sprintf(psqlExecuteSP, procedure))
	vs := make([]interface{}, 0)
	var cont int
	cont = 1
	_, err := r.WriteString("( ")

	if err != nil {
		logger.Error.Printf("agregando parametro usuario a la ejecucion del SP en psqlExecuteSP: %v", err)
		return rs, err
	}

	for i, v := range parameters {
		if len(v.(string)) > 0 {
			_, err := r.WriteString(fmt.Sprintf(`%s => $%d,`, i, cont))
			if err != nil {
				logger.Error.Printf("agregando parametros a la ejecucion del SP en psqlExecuteSP: %v", err)
				return rs, err
			}
			vs = append(vs, v)
			cont++
		}
	}
	if option == 1 {

		_, err := r.WriteString(fmt.Sprintf(`user_id => $%d,`, cont))
		if err != nil {
			logger.Error.Printf("agregando parametro usuario a la ejecucion del SP en sqlserverExecuteSP: %v", err)
			return rs, err
		}

		vs = append(vs, s.user.ID)
		cont++
	}
	r.Truncate(r.Len() - 1)
	_, err = r.WriteString(" )")
	if err != nil {
		logger.Error.Printf("agregando cierre query executeSP: %v", err)
		return rs, err
	}
	stmt, err := s.DB.Prepare(r.String())
	if err != nil {
		logger.Error.Printf("preparando consulta psqlExecuteSP: %s, %v", r.String(), err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(vs...)
	if err != nil {
		logger.Error.Printf("ejecutando psqlExecuteSP user: %s, %v ", r.String(), err)
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
			logger.Error.Printf("no se pudo escanear las columnas de la consulta psqlGetInfoktg: %t", err)
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

func (s *psql) ExecuteSPBYDocumentID(procedure string, documentID int64) (int, error) {
	const psqlExecuteSP = `select * from %s( $1, $2)`
	var res int
	sqlExecute := fmt.Sprintf(psqlExecuteSP, procedure)
	stmt, err := s.DB.Prepare(sqlExecute)
	if err != nil {
		logger.Error.Printf("preparando la sentencia ExecuteSP: %v", err)
		return 0, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		&documentID,
		&s.user.ID,
	).Scan(&res)
	if err != nil {
		logger.Error.Printf("***ejecutando la sentencia ExecuteSP: %v", err)
		return 0, err
	}
	return res, nil
}

func NewReportPsqlRepository(db *sqlx.DB, user *models.User, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}
