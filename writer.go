package main

import (
	"encoding/csv"
	"os"

	bf "gopkg.in/russross/blackfriday.v2"
)

const (
	HTML Option = 1 << iota
)

type Option int

type DeckWriter struct {
	fp      *os.File
	options Option
}

// NewDeckWriter return the DeckWriter struct which contain a filepointer `fp`
// to which the Deck can be written to.
func NewDeckWriter(fp *os.File, options Option) *DeckWriter {
	return &DeckWriter{fp: fp, options: options}
}

// WriteDeck will write the Deck to the filepointer `fp` in a specified format
func (dw *DeckWriter) WriteDeck(d *Deck) error {
	err := dw.writeToCSV(d)
	if err != nil {
		return err
	}

	return nil
}

// writeToCSV will write the Deck (AST) in the specified csv format
func (dw *DeckWriter) writeToCSV(d *Deck) error {
	csvWriter := csv.NewWriter(dw.fp)
	defer csvWriter.Flush()

	for _, c := range d.Cards {
		var row []string

		for _, f := range c.Fields {
			var c string
			switch dw.options {
			case HTML:
				c = string(bf.Run([]byte(f.Content)))
			default:
				c = f.Content
			}
			row = append(row, c)
		}

		csvWriter.Write(row)
	}

	return nil
}
