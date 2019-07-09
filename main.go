package main

import (
	"errors"
	"flag"
	"github.com/Brialius/gocopy/gocopy"
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

var (
	args    = Args{}
	Version = "dev"
	Build   = "local"
)

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
	}

	if args.Verbosity {
		log.Printf("from: %s (size: %d), to: %s, offset: %d, limit: %d",
			args.From, inputSize, args.To, args.Offset, args.Limit)
	}
	return nil
}

func main() {
	log.Printf("gocopy %s-%s", Version, Build)
	flag.Parse()
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
