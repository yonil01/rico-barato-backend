package upload_metadata

import "time"

type Metadata struct {
	ID                          *string   `json:"id" db:"id"`
	DocumentID                  *string   `json:"document_id" db:"document_id"`
	NumeroCarpeta               *string   `json:"numero_carpeta" db:"numero_carpeta"`
	NumeroContrato              *string   `json:"numero_contrato" db:"numero_contrato"`
	RazonSocial                 *string   `json:"razon_social" db:"razon_social"`
	Etapa                       *string   `json:"etapa" db:"etapa"`
	Tomo                        *string   `json:"tomo" db:"tomo"`
	Serie                       *string   `json:"serie" db:"serie"`
	NumeroLote                  *string   `json:"numero_lote" db:"numero_lote"`
	ObjetoContractual           *string   `json:"objeto_contractual" db:"objeto_contractual"`
	TipologiaDocumental         *string   `json:"tipologia_documental" db:"tipologia_documental"`
	NumeroRadicado              *string   `json:"numero_radicado" db:"numero_radicado"`
	NumeroDelDocumento          *string   `json:"numero_del_documento" db:"numero_del_documento"`
	Idioma                      *string   `json:"idioma" db:"idioma"`
	GlobalId                    *string   `json:"global_id" db:"global_id"`
	DpiPpp                      *string   `json:"dpi_ppp" db:"dpi_ppp"`
	PaginasElectronicas         *string   `json:"paginas_electronicas" db:"paginas_electronicas"`
	SoftwareCaptura             *string   `json:"software_captura" db:"software_captura"`
	HardwareCaptura             *string   `json:"hardware_captura" db:"hardware_captura"`
	FechaUltimaCalibracion      *string   `json:"fecha_ultima_calibracion" db:"fecha_ultima_calibracion"`
	FechaFirma                  *string   `json:"fecha_firma" db:"fecha_firma"`
	ActaDeDigitalizacionInicial *string   `json:"acta_de_digitalizacion_inicial" db:"acta_de_digitalizacion_inicial"`
	FechaCreacionActa           *string   `json:"fecha_creacion_acta" db:"fecha_creacion_acta"`
	ActaDeDigitalizacionFinal   *string   `json:"acta_de_digitalizacion_final" db:"acta_de_digitalizacion_final"`
	NumeroLotes                 *string   `json:"numero_lotes" db:"numero_lotes"`
	FechaSello                  *string   `json:"fecha_sello" db:"fecha_sello"`
	Peso                        *string   `json:"peso" db:"peso"`
	RutaImagen                  *string   `json:"ruta_imagen" db:"ruta_imagen"`
	FechaDelDocumento           *string   `json:"fecha_del_documento" db:"fecha_del_documento"`
	FechaDigitalizacion         *string   `json:"fecha_digitalizacion" db:"fecha_digitalizacion"`
	Fuente                      *string   `json:"fuente" db:"fuente"`
	NombreArchivo               *string   `json:"nombre_archivo" db:"nombre_archivo"`
	CreatedAt                   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt                   time.Time `json:"updated_at" db:"updated_at"`
}

func NewMetadata(ID *string, DocumentID *string, RazonSocial *string,
	Etapa *string, Tomo *string, Serie *string, NumeroLote *string, ObjetoContractual *string, TipologiaDocumental *string,
	NumeroRadicado *string, NumeroDelDocumento *string, Idioma *string, GlobalId *string, DpiPpp *string,
	PaginasElectronicas *string, SoftwareCaptura *string, HardwareCaptura *string, FechaUltimaCalibracion *string,
	FechaFirma *string, ActaDeDigitalizacionInicial *string, FechaCreacionActa *string, ActaDeDigitalizacionFinal *string,
	NumeroLotes *string, FechaSello *string, FechaDelDocumento *string,
	FechaDigitalizacion *string, StatusInventarioIbml *string, Fuente *string) *Metadata {
	return &Metadata{
		ID:                          ID,
		DocumentID:                  DocumentID,
		RazonSocial:                 RazonSocial,
		Etapa:                       Etapa,
		Tomo:                        Tomo,
		Serie:                       Serie,
		NumeroLote:                  NumeroLote,
		ObjetoContractual:           ObjetoContractual,
		TipologiaDocumental:         TipologiaDocumental,
		NumeroRadicado:              NumeroRadicado,
		NumeroDelDocumento:          NumeroDelDocumento,
		Idioma:                      Idioma,
		GlobalId:                    GlobalId,
		DpiPpp:                      DpiPpp,
		PaginasElectronicas:         PaginasElectronicas,
		SoftwareCaptura:             SoftwareCaptura,
		HardwareCaptura:             HardwareCaptura,
		FechaUltimaCalibracion:      FechaUltimaCalibracion,
		FechaFirma:                  FechaFirma,
		ActaDeDigitalizacionInicial: ActaDeDigitalizacionInicial,
		FechaCreacionActa:           FechaCreacionActa,
		ActaDeDigitalizacionFinal:   ActaDeDigitalizacionFinal,
		NumeroLotes:                 NumeroLotes,
		FechaSello:                  FechaSello,
		FechaDelDocumento:           FechaDelDocumento,
		FechaDigitalizacion:         FechaDigitalizacion,
		Fuente:                      Fuente,
	}
}
