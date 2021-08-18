package bytechomp

import (
	"debug/elf"
	"encoding/base64"
	"encoding/csv"
	"os"
)

var (
	elfFile *ElfFile
)

//ElfFile a struct representing all of the information contained within an ELF binary object
type ElfFile struct {
	file        *elf.File
	sectionData map[sectionOfInterest][]byte
}

//NewElfFile creates a new ElfFile struct, should be used instead of new
func NewElfFile(file *elf.File) *ElfFile {
	elfFile = &ElfFile{file, make(map[sectionOfInterest][]byte)}
	return elfFile
}

//StartAnalysis starts analysis on an ELF binary, returns an error if a problem occurs
func (elfFile *ElfFile) StartAnalysis() error {
	err := elfFile.loadSections()
	if err != nil {
		return err
	}
	return nil
}

func (elfFile *ElfFile) loadSectionData(sectionType sectionOfInterest) error {
	section := elfFile.file.Section(string(sectionType))
	if data, err := section.Data(); err != nil {
		return err
	} else {
		elfFile.sectionData[sectionType] = data
	}
	return nil
}

func (elfFile *ElfFile) loadSections() error {
	for _, section := range sectionsOfInterest {
		err := elfFile.loadSectionData(section)
		if err != nil {
			return err
		}
	}
	return nil
}

func (elfFile *ElfFile) getBase64() map[string]*string {
	var base64Map = make(map[string]*string, len(elfFile.sectionData))
	for key, value := range elfFile.sectionData {
		encodedString := base64.StdEncoding.EncodeToString(value)
		base64Map[string(key)] = &encodedString
	}
	return base64Map
}

func (elfFile *ElfFile) CreateCSV() (fileName string, err error) {
	file, err := os.Create("result.csv")
	if err != nil {
		return "", err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
}
