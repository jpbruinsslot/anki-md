package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	VERSION = "0.1.2"
	USAGE   = `NAME:
    anki-md - markdown to anki flash cards converter

USAGE:
    anki-md -i [input-file] -o [output-file]

EXAMPLES:

    $ anki-md -i deck.md -o deck.csv

    $ cat deck.md | anki-md -o deck.csv

    $ anki-md -i deck.md > test.csv

VERSION:
    %s

WEBSITE:
    https://github.com/jpbruinsslot/anki-md

GLOBAL OPTIONS:
    -i, -input [input-file]     input file
    -o, -output [output-file]   output file
    -html                       convert field content to html
    -h, -help
`
)

var (
	flgInput  string
	flgOutput string
	flgHTML   bool
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

	flag.BoolVar(
		&flgHTML,
		"html",
		false,
		"convert field content to html",
	)

	flag.Usage = func() {
		fmt.Printf(USAGE, VERSION)
	}
}

func main() {
	flag.Parse()

	var err error

	// Input
	var r io.Reader
	if flgInput != "" {
		r, err = os.Open(flgInput)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		r = os.Stdin
	}

	// Output
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

	// Flag options
	var options Option
	if flgHTML {
		options = options | HTML
	}

	// Parse input
	d, err := NewParser(r).ParseDeck()
	if err != nil {
		log.Fatal(err)
	}

	// Write output
	err = NewDeckWriter(fp, options).WriteDeck(d)
	if err != nil {
		log.Fatal(err)
	}
}
