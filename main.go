package main

import (
	"errors"
	"flag"
	"github.com/Brialius/gocopy/gocopy"
	"log"
	"os"
)

// Args is a structure to store command line parsed parameters
type Args struct {
	From      string
	To        string
	Offset    int64
	Limit     int64
	Verbosity bool
}

var (
	args = Args{}
	// Version is a application version substituted from CI build
	Version = "dev"
	// Build is a git hash substituted from CI build
	Build = "local"
)

func init() {
	log.SetFlags(0)
	flag.StringVar(&args.From, "from", "", "file to read from")
	flag.StringVar(&args.To, "to", "", "file to write to")
	flag.Int64Var(&args.Offset, "offset", 0, "offset in input file (should be >= 0)")
	flag.Int64Var(&args.Limit, "limit", -1, "block size to copy")
	flag.BoolVar(&args.Verbosity, "v", false, "verbosity mode")
}

func validateArgs(a Args) error {
	var inputSize int64
	if len(a.From) == 0 || len(a.To) == 0 {
		return errors.New("input or output file is not defined")
	}

	if a.Verbosity {
		log.Printf("from: %s (size: %d), to: %s, offset: %d, limit: %d",
			a.From, inputSize, a.To, a.Offset, a.Limit)
	}
	return nil
}

func main() {
	log.Printf("gocopy %s-%s", Version, Build)
	flag.Parse()
	if err := validateArgs(args); err != nil {
		if args.Verbosity {
			log.Printf("args: %#v\n", args)
		}
		flag.Usage()
		log.Fatalf("Args validation error: %s\n", err)
	}

	if err := gocopy.Copy(args.From, args.To, args.Offset, args.Limit); err != nil {
		log.Printf("Copy error: %s", err)
		os.Exit(1)
	}
}
