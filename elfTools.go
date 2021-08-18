package bytechomp

import (
	"debug/elf"
	"os"
)

var elfFile *ElfFile

type ElfFile struct {
	file       *os.File
	fileHeader *elf.FileHeader
	header32   *elf.Header32
	header64   *elf.Header64
}

//LoadElf loads an ELF binary to be analyzed, returns nil if loaded successfully
func LoadElf(file *os.File) error {
	if _, err := IsElf(file); err != nil {
		return err
	}
	elfFile = &ElfFile{}
	return nil
}
