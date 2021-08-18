package bytechomp

import (
	"debug/elf"
)

var (
	elfFile *ElfFile
)

//ElfFile a struct representing all of the information contained within an ELF binary object
type ElfFile struct {
	file        *elf.File
	projectName string
	sectionData map[sectionOfInterest][]byte
}

//NewElfFile creates a new ElfFile struct, should be used instead of new
func NewElfFile(file *elf.File, projectName string) *ElfFile {
	elfFile = &ElfFile{file, projectName, make(map[sectionOfInterest][]byte)}
	return elfFile
}

//GenerateCSV() generates a csv of ELF file information
func (elfFile *ElfFile) GenerateCSV() {
	if len(elfFile.sectionData) == 0 {
		elfFile.StartAnalysis()
	}
	csvData := MakeCsvData(elfFile)
	csvData.CreateCSV()
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
