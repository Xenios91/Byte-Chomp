package bytechomp

import (
	"debug/elf"
	"os"
	"testing"
)

//useless test, just using for testing while coding
func TestCreateCsv(t *testing.T) {
	testGoBin := "./testGoBin"
	testCSV := "./testProject.csv"
	f, _ := elf.Open(testGoBin)
	elfFile := NewElfFile(f, "testProject")
	elfFile.GenerateCSV()
	if _, err := os.Stat(testCSV); os.IsNotExist(err) {
		t.FailNow()
	} else {
		os.Remove(testCSV)
	}
}
