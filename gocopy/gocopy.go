package gocopy

import (
	"errors"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"io"
	"log"
	"os"
)

func Copy(srcPath, dstPath string, offset, limit int64) error {
	var src, dst *os.File
	var inputSize int64

	if stat, err := os.Stat(srcPath); err == nil {
		inputSize = stat.Size()
	} else {
		return fmt.Errorf("input file can't be opened: %s", err)
	}

	if limit == -1 {
		limit = inputSize - offset
	}

	if err := validate(limit, inputSize, offset); err != nil {
		return err
	}

	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}

	defer func() {
		if err := src.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	if _, err := src.Seek(offset, io.SeekStart); err != nil {
		return err
	}

	dst, err = os.Create(dstPath)
	if err != nil {
		return err
	}

	buf := make([]byte, 1024*64) // 64kb
	var written int64
	err = nil
	bar := pb.StartNew(int(limit))
	bar.Set(pb.Bytes, true)
	bar.Start()

	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
				bar.Add(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}

	bar.Finish()

	if err := dst.Close(); err != nil {
		log.Fatalln(err)
	}

	return err
}

func validate(limit, inputSize, offset int64) error {
	switch {
	case inputSize == 0:
		return errors.New("input file is empty")
	case limit <= 0:
		return fmt.Errorf("limit parameter is: %d, but should be > 0", limit)
	case offset < 0:
		return fmt.Errorf("offset parameter is: %d, but should be >= 0", offset)
	case offset >= inputSize:
		return fmt.Errorf("offset (%d) is greater or equal input file size (%d)", offset, inputSize)
	case (limit + offset) > inputSize:
		return fmt.Errorf("limit + offset is greater sorurce size: %d > %d", limit+offset, inputSize)
	}
	return nil
}
