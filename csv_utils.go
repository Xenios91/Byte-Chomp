package bytechomp

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"os"
)

//CsvData a struct representing the data to be exported to a csv file
type CsvData struct {
	fileName string
	headers  []string
	rows     []string
}

//MakeCsvData creates a new CsvData struct and fills out some of its attributes based on the elfFile struct passed to it
func MakeCsvData(elfFile *ElfFile) *CsvData {
	size := len(elfFile.sectionData)
	csvData := &CsvData{elfFile.projectName, make([]string, size), make([]string, size)}
	return csvData
}

func (csvData *CsvData) getBase64() map[string]*string {
	var base64Map = make(map[string]*string, len(elfFile.sectionData))
	for key, value := range elfFile.sectionData {
		encodedString := base64.StdEncoding.EncodeToString(value)
		base64Map[string(key)] = &encodedString
	}
	return base64Map
}

func (csvData *CsvData) getCSVSectionHeaders() []string {
	var headers = make([]string, len(elfFile.sectionData))
	var counter = 0

	for key := range elfFile.sectionData {
		headers[counter] = string(key)
		counter++
	}
	return headers
}

func (csvData *CsvData) getRows() []string {
	var data = make([]string, len(elfFile.sectionData))
	base64Data := csvData.getBase64()
	var counter int = 0

	for _, value := range csvData.headers {
		data[counter] = *base64Data[value]
		counter++
	}
	return data
}

//CreateCSV creates a new csv file from the CsvData struct that calls this method
func (csvData *CsvData) CreateCSV() (fileName string, err error) {
	fileName = fmt.Sprintf("%s.csv", csvData.fileName)
	file, err := os.Create(fileName)
	if err != nil {
		return fileName, err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	csvData.headers = csvData.getCSVSectionHeaders()
	csvData.rows = csvData.getRows()

	writer.Write(csvData.headers)
	writer.Write(csvData.rows)

	return fileName, err
}
