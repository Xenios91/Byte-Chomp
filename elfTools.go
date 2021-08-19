package bytechomp

import (
	"debug/elf"
	"errors"
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
func NewElfFile(file *elf.File, projectName string) (*ElfFile, error) {
	if file == nil {
		return nil, errors.New("elf file cannot be nil")
	}
	elfFile = &ElfFile{file, projectName, make(map[sectionOfInterest][]byte)}
	return elfFile, nil
}

func (elfFile *ElfFile) validateElfFile() error {
	if elfFile.file == nil {
		return errors.New("invalid elf file")
	}
	return nil
}

//GenerateCSV generates a csv of ELF file information
func (elfFile *ElfFile) GenerateCSV() (fileName string, err error) {
	if len(elfFile.sectionData) == 0 {
		elfFile.StartAnalysis()
	}
	csvData := MakeCsvData(elfFile)
	fileName, err = csvData.CreateCSV()
	return fileName, err
}

//StartAnalysis starts analysis on an ELF binary, returns an error if a problem occurs
func (elfFile *ElfFile) StartAnalysis() error {
	if err := elfFile.validateElfFile(); err != nil {
		return errors.New("elf file failed validations")
	}
	if err := elfFile.loadSections(); err != nil {
		return errors.New("sections of the elf file could not be loaded")
	}
	return nil
}

func (elfFile *ElfFile) loadSectionData(sectionType sectionOfInterest) error {
	section := elfFile.file.Section(string(sectionType))
	data, err := section.Data()
	if err != nil {
		return err
	}
	elfFile.sectionData[sectionType] = data
	return nil
}

func (elfFile *ElfFile) loadSections() error {
	var err error
	for _, section := range sectionsOfInterest {
		err = elfFile.loadSectionData(section)
		if err != nil {
			break
		}
	}
	return err
}
