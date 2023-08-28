package generator

import (
	"backend-ccff/internal/env"
	"backend-ccff/internal/logger"
	"backend-ccff/pkg/indra/upload_metadata"
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"strconv"
	"time"
)

func GeneratorXLSX(data []*upload_metadata.Metadata) (string, error) {

	startIndex := env.NewConfiguration().App.IndexSeparator
	divideIndex := len(data) / startIndex
	f := excelize.NewFile()
	for i := 0; i <= divideIndex; i++ {
		f.NewSheet(fmt.Sprintf("Sheet%d", i+1))
		// Generar los encabezados
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "A1", "id")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "B1", "document_id")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "C1", "numero_carpeta")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "D1", "numero_contrato")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "E1", "razon_social")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "F1", "etapa")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "G1", "tomo")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "H1", "serie")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "I1", "numero_lote")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "J1", "objeto_contractual")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "K1", "tipologia_documental")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "L1", "numero_radicado")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "M1", "numero_del_documento")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "N1", "idioma")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "O1", "global_id")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "P1", "dpi_ppp")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "Q1", "paginas_electronicas")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "R1", "software_captura")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "S1", "hardware_captura")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "T1", "fecha_ultima_calibracion")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "U1", "fecha_firma")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "V1", "acta_de_digitalizacion_inicial")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "W1", "fecha_creacion_acta")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "X1", "acta_de_digitalizacion_final")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "Y1", "numero_lotes")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "Z1", "fecha_sello")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "AA1", "peso")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "AB1", "ruta_imagen")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "AC1", "fecha_del_documento")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "AD1", "fecha_digitalizacion")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "AF1", "fuente")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "AG1", "nombre_archivo")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "AH1", "created_at")
		f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), "AI1", "updated_at")

		indexStart := startIndex * i
		endStart := startIndex * (i + 1)
		if endStart > len(data) {
			endStart = len(data)
		}
		indexRow := 0
		for index := indexStart; index < endStart; index++ {
			row := strconv.Itoa(indexRow + 2)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("A%s", row), *data[index].ID)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("B%s", row), *data[index].DocumentID)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("C%s", row), *data[index].NumeroCarpeta)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("D%s", row), *data[index].NumeroContrato)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("E%s", row), *data[index].RazonSocial)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("F%s", row), *data[index].Etapa)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("G%s", row), *data[index].Tomo)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("H%s", row), *data[index].Serie)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("I%s", row), *data[index].NumeroLote)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("J%s", row), *data[index].ObjetoContractual)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("K%s", row), *data[index].TipologiaDocumental)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("L%s", row), *data[index].NumeroRadicado)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("M%s", row), *data[index].NumeroDelDocumento)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("N%s", row), *data[index].Idioma)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("O%s", row), *data[index].GlobalId)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("P%s", row), *data[index].DpiPpp)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("Q%s", row), *data[index].PaginasElectronicas)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("R%s", row), *data[index].SoftwareCaptura)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("S%s", row), *data[index].HardwareCaptura)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("T%s", row), *data[index].FechaUltimaCalibracion)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("U%s", row), *data[index].FechaFirma)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("V%s", row), *data[index].ActaDeDigitalizacionInicial)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("W%s", row), *data[index].FechaCreacionActa)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("X%s", row), *data[index].ActaDeDigitalizacionFinal)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("Y%s", row), *data[index].NumeroLotes)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("Z%s", row), *data[index].FechaSello)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("AA%s", row), *data[index].Peso)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("AB%s", row), *data[index].RutaImagen)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("AC%s", row), *data[index].FechaDelDocumento)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("AD%s", row), *data[index].FechaDigitalizacion)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("AF%s", row), *data[index].Fuente)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("AG%s", row), *data[index].NombreArchivo)
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("AH%s", row), data[index].CreatedAt.String())
			f.SetCellValue(fmt.Sprintf("Sheet%d", i+1), fmt.Sprintf("AI%s", row), data[index].UpdatedAt.String())
			indexRow++
		}
	}

	pathDirectory := env.NewConfiguration().App.PathDirectory
	nameFolder := time.Now().Format("2006-01-02")
	namePath := fmt.Sprintf("%s/%s", pathDirectory, nameFolder)
	err := createPathFolder(namePath)
	if err != nil {
		logger.Error.Printf("Couldn't createPathFolder: %v", err)
		return "", err
	}

	nameFile := fmt.Sprintf("METADATA_SOLICITADA_%d_%d_%d.xlsx", time.Now().Day(), time.Now().Minute(), time.Now().Second())
	pathFile := fmt.Sprintf("%s/%s", namePath, nameFile)
	if err := f.SaveAs(pathFile); err != nil {
		return "", err
	}

	return pathFile, nil
}

func createPathFolder(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			logger.Error.Printf("Couldn't MkdirAll: %v", err)
			return err
		}
	}
	return nil
}
