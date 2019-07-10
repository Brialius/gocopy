package gocopy

import (
	"errors"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"io"
	"log"
	"os"
)

// Copy is a function to copy section of file (or whole file) to another place with progress bar
func Copy(srcPath, dstPath string, offset, limit int64) error {
	var (
		src, dst  *os.File
		inputSize int64
	)

	if stat, err := os.Stat(srcPath); err != nil {
		return fmt.Errorf("input file can't be opened: %s", err)
	} else {
		inputSize = stat.Size()
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

	dst, err = os.Create(dstPath)
	if err != nil {
		return err
	}

	_, errRes := copyNWithOffset(src, dst, limit, offset)

	if err := dst.Close(); err != nil {
		log.Fatalln(err)
	}
	return errRes
}

func copyNWithOffset(src io.ReaderAt, dst io.Writer, n, offset int64) (int64, error) {
	var (
		written int64
		errRes  error
		buf     []byte
	)
	// 64kb
	const bufSize = 1024 * 64

	if n < bufSize {
		buf = make([]byte, n)
	} else {
		buf = make([]byte, bufSize)
	}

	bar := pb.StartNew(int(n))
	bar.Set(pb.Bytes, true)
	bar.Start()

	sr := io.NewSectionReader(src, offset, n)

	for {
		nr, er := sr.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
				bar.Add(nw)
			}
			if ew != nil {
				errRes = ew
				break
			}
			if nr != nw {
				errRes = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				errRes = er
			}
			break
		}
	}

	bar.Finish()
	return written, errRes
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
