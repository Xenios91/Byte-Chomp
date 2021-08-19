package bytechomp

import (
	"debug/elf"
	"os"
	"testing"
)

func getTestGoBinFile() (*ElfFile, error) {
	testGoBin := "./testGoBin"
	file, _ := os.Open(testGoBin)
	elf, _ := elf.NewFile(file)
	elfFile, err := NewElfFile(elf, "testProjectName")
	return elfFile, err
}

//useless test, just using for testing while coding
func TestCreateCsv(t *testing.T) {
	elfFile, _ := getTestGoBinFile()
	fileName, err := elfFile.GenerateCSV()
	if err != nil {
		t.FailNow()
	}
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		t.FailNow()
	} else {
		os.Remove(fileName)
	}
}

func TestNewElfFile(t *testing.T) {
	testGoBin := "./testGoBin"
	file, err := os.Open(testGoBin)
	if err != nil {
		t.FailNow()
	}

	elfNativeFile, err := elf.NewFile(file)
	if err != nil {
		t.FailNow()
	}

	_, err = NewElfFile(elfNativeFile, "testProjectName")
	if err != nil {
		t.FailNow()
	}

	var fileDoesntExist *elf.File = nil
	_, err = NewElfFile(fileDoesntExist, "testBadProject")
	if err == nil {
		t.FailNow()
	}
}

func TestStartAnalysis(t *testing.T) {
	elfFile, err := getTestGoBinFile()
	if err != nil {
		t.FailNow()
	}
	err = elfFile.StartAnalysis()
	if err != nil {
		t.FailNow()
	}

	elfFile.file = nil
	err = elfFile.StartAnalysis()
	if err == nil {
		t.FailNow()
	}
}
