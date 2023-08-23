package dni

type ResponseReniec struct {
	Success         bool    `json:"success" db:"success"`
	Dni             *string `json:"dni" db:"dni"`
	Nombres         *string `json:"nombres" db:"nombres"`
	ApellidoPaterno *string `json:"apellidoPaterno" db:"apellidoPaterno"`
	ApellidoMaterno *string `json:"apellidoMaterno" db:"apellidoMaterno"`
	CodVerifica     *int    `json:"codVerifica" db:"codVerifica"`
}
