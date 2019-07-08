package gocopy

import (
	"io"
	"os"
)

func Copy(srcPath, dstPath string, offset, limit int64) error {
	var src, dst *os.File

	src, err := os.Open(srcPath)

	if err != nil {
		return err
	}
	_, _ = src.Seek(offset, io.SeekStart)
	defer src.Close()

	dst, err = os.Create(dstPath)
	if err != nil {
		return err
	}

	if _, err := io.CopyN(dst, src, limit); err != nil && err != io.EOF {
		return err
	}
	_ = dst.Close()
	return nil
}
