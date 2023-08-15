package dni

type ResponseReniec struct {
	Success *bool      `json:"success" db:"success"`
	Data    DataReniec `json:"data" db:"data"`
}

type DataReniec struct {
	Numero          *string      `json:"numero" db:"numero"`
	NombreCompleto  *string      `json:"nombre_completo" db:"nombre_completo"`
	Nombres         *string      `json:"nombres" db:"nombres"`
	ApellidoPaterno *string      `json:"apellido_paterno" db:"apellido_paterno"`
	ApellidoMaterno *string      `json:"apellido_materno" db:"apellido_materno"`
	UbigeoSunat     *string      `json:"ubigeo_sunat" db:"ubigeo_sunat"`
	Ubigeo          *interface{} `json:"ubigeo" db:"ubigeo"`
}
