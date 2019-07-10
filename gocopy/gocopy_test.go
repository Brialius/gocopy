package gocopy

import (
	"bytes"
	"io"
	"testing"
)

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

func Test_copyNWithOffset(t *testing.T) {
	type args struct {
		src    io.ReaderAt
		n      int64
		offset int64
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantDst string
		wantErr bool
	}{
		{
			"OK full copy",
			args{bytes.NewReader([]byte("12345")), 5, 0},
			5,
			"12345",
			false,
		},
		{
			"OK copy offset",
			args{bytes.NewReader([]byte("12345")), 4, 1},
			4,
			"2345",
			false,
		},
		{
			"OK copy offset",
			args{bytes.NewReader([]byte("12345")), 3, 1},
			3,
			"234",
			false,
		},
		{
			"OK copy limit",
			args{bytes.NewReader([]byte("12345")), 3, 0},
			3,
			"123",
			false,
		},
		{
			"Fail copy negative offset",
			args{bytes.NewReader([]byte{}), 3, -1},
			0,
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dst := &bytes.Buffer{}
			got, err := copyNWithOffset(tt.args.src, dst, tt.args.n, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("copyNWithOffset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("copyNWithOffset() = %v, want %v", got, tt.want)
			}
			if gotDst := dst.String(); gotDst != tt.wantDst {
				t.Errorf("copyNWithOffset() dst = %v, want %v", gotDst, tt.wantDst)
			}
		})
	}
}
