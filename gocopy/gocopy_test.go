package gocopy

import "testing"

func TestCopy(t *testing.T) {
	type args struct {
		srcPath string
		dstPath string
		offset  int64
		limit   int64
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
			if err := Copy(tt.args.srcPath, tt.args.dstPath, tt.args.offset, tt.args.limit); (err != nil) != tt.wantErr {
				t.Errorf("Copy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
