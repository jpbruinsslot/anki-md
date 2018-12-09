package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	VERSION = "0.1.0"
	USAGE   = `NAME:
    anki-md - markdown to anki flash cards converter

USAGE:
    anki-md -i [input-file] -o [output-file]

VERSION:
    %s

WEBSITE:
    https://github.com/erroneousboat/anki-md

GLOBAL OPTIONS:
    -i, -input [input-file]
    -o, -output [output-file]
    -h, -help
`
)

var (
	flgInput  string
	flgOutput string
)

func init() {
	flag.StringVar(
		&flgInput,
		"i",
		"",
		"anki-md input file",
	)

	flag.StringVar(
		&flgInput,
		"input",
		"",
		"anki-md input file",
	)

	flag.StringVar(
		&flgOutput,
		"o",
		"",
		"output file",
	)

	flag.StringVar(
		&flgOutput,
		"output",
		"",
		"output file",
	)

	flag.Usage = func() {
		fmt.Printf(USAGE, VERSION)
	}

	flag.Parse()
}

func main() {
	var err error

	var r io.Reader
	if flgInput != "" {
		r, err = os.Open(flgInput)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		r = os.Stdin
	}

	var fp *os.File
	if flgOutput != "" {
		fp, err = os.Create(flgOutput)
		if err != nil {
			log.Fatal(err)
		}
		defer fp.Close()
	} else {
		fp = os.Stdout
	}

	d, err := NewParser(r).ParseDeck()
	if err != nil {
		log.Fatal(err)
	}

	err = NewDeckWriter(fp).WriteDeck(d)
	if err != nil {
		log.Fatal(err)
	}
}
