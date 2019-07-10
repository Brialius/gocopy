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

func Test_validate(t *testing.T) {
	type args struct {
		limit     int64
		inputSize int64
		offset    int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"OK",
			args{1, 5, 0},
			false,
		}, {
			"Fail empty source",
			args{1, 0, 0},
			true,
		}, {
			"Fail negative limit",
			args{-10, 5, 0},
			true,
		}, {
			"Fail zero limit",
			args{0, 5, 0},
			true,
		}, {
			"Fail negative offset",
			args{1, 5, -1},
			true,
		}, {
			"Fail limit + offset is greater sorurce size",
			args{5, 5, 1},
			true,
		}, {
			"OK limit + offset is equal sorurce size",
			args{5, 6, 1},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate(tt.args.limit, tt.args.inputSize, tt.args.offset); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
