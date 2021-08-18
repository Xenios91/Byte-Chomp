package bytechomp

import (
	"os"
)

//IsElf checks whether an binary object is of type ELF and returns a boolean
func IsElf(file *os.File) (bool, error) {
	var ident [16]uint8

	if _, err := file.ReadAt(ident[0:], 0); err != nil {
		return false, err
	} else if ident[0] != '\x7f' || ident[1] != 'E' || ident[2] != 'L' || ident[3] != 'F' {
		return false, nil
	}
	return true, nil
}
