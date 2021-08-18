package bytechomp

import (
	"debug/elf"
	"testing"
)

//useless test, just using for testing while coding
func TestThingy(t *testing.T) {
	f, _ := elf.Open("./test")
	elfFile := NewElfFile(f)
	elfFile.StartAnalysis()
}
