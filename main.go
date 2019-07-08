package main

import (
	"errors"
	"flag"
	"fmt"
	"gocopy/gocopy"
	"log"
	"os"
)

type Args struct {
	From      string
	To        string
	Offset    int64
	Limit     int64
	Verbosity bool
}

var args = Args{}
var Version = "dev"
var Build = "undefined"

func init() {
	log.SetFlags(0)
	flag.StringVar(&args.From, "from", "", "file to read from")
	flag.StringVar(&args.To, "to", "", "file to write to")
	flag.Int64Var(&args.Offset, "offset", 0, "offset in input file (should be >= 0)")
	flag.Int64Var(&args.Limit, "limit", -1, "block size to copy")
	flag.BoolVar(&args.Verbosity, "v", false, "verbosity mode")
}

func validateArgs() error {
	var inputSize int64

	if len(args.From) == 0 || len(args.To) == 0 {
		return errors.New("input or output file is not defined")
	} else {
		if stat, err := os.Stat(args.From); err == nil {
			inputSize = stat.Size()
		} else {
			return fmt.Errorf("input file can't be opened: %s", err)
		}
	}

	if args.Offset < 0 {
		return fmt.Errorf("offset parameter is (%d), but should be >= 0", args.Offset)
	}

	if args.Offset >= inputSize {
		return fmt.Errorf("offset (%d) is greater or equal input file size (%d)\n", args.Offset, inputSize)
	}

	if args.Limit == 0 {
		return fmt.Errorf("limit parameter is (%d), but should be > 0", args.Limit)
	}

	if args.Limit == -1 {
		args.Limit = inputSize
	}

	if args.Verbosity {
		log.Printf("from: %s (size: %d), to: %s, offset: %d, limit: %d",
			args.From, inputSize, args.To, args.Offset, args.Limit)
	}
	return nil
}

func main() {
	flag.Parse()
	log.Printf("gocopy %s-%s", Version, Build)
	if err := validateArgs(); err != nil {
		log.Printf("Args validation error: %s\n", err)
		if args.Verbosity {
			log.Printf("args: %#v\n", args)
		}
		flag.Usage()
		os.Exit(1)
	}

	if err := gocopy.Copy(args.From, args.To, args.Offset, args.Limit); err != nil {
		log.Printf("Copy error: %s", err)
		os.Exit(1)
	}
}
