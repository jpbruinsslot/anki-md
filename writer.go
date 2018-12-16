package main

import (
	"encoding/csv"
	"os"
)

type DeckWriter struct {
	fp *os.File
}

// NewDeckWriter return the DeckWriter struct which contain a filepointer `fp`
// to which the Deck can be written to.
func NewDeckWriter(fp *os.File) *DeckWriter {
	return &DeckWriter{fp: fp}
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
			row = append(row, f.Content)
		}
		csvWriter.Write(row)
	}

	return nil
}
