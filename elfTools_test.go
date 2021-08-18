package bytechomp

import (
	"os"
	"testing"
)

func TestLoadElf(t *testing.T) {
	type args struct {
		file *os.File
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoadElf(tt.args.file); (err != nil) != tt.wantErr {
				t.Errorf("LoadElf() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
