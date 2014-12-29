package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	skip := flag.Int64("skip", 0, "Skip bytes at start")
	lim := flag.Int64("lim", 0, "Limit read bytes")
	flag.Usage = usage
	flag.Parse()

	var r io.Reader

	if n := flag.NArg(); n == 0 || n == 1 && flag.Arg(0) == "-" {
		// stdin
		r = os.Stdin
	} else if flag.NArg() == 1 {
		f, err := os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		r = f
	} else {
		usage()
		os.Exit(2) // same as flag package when there are invalid flags
	}

	// If a skip was set, seek ahead that many bytes
	if *skip != 0 {
		rs := r.(io.ReadSeeker) // both files and stdin are seekable
		n, err := rs.Seek(*skip, os.SEEK_SET)
		if err != nil {
			log.Fatal(err)
		}
		if n != *skip {
			log.Fatalf("Skipped %d != %d", n, *skip)
		}
	}

	// If a limit was set, stop reading after that many bytes
	if *lim != 0 {
		r = io.LimitReader(r, *lim)
	}

	dumpReader(r)
}

func dumpReader(r io.Reader) {
	d := hex.Dumper(os.Stdout)

	_, err := io.Copy(d, r)
	if err != nil {
		log.Fatal(err)
	}

	err = d.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func usage() {
	fmt.Printf("Usage:\n  hd [options] [filename]\n\nOptions:\n")
	flag.PrintDefaults()
}
