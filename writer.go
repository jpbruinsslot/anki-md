package main

import (
	"encoding/csv"
	"os"
)

type DeckWriter struct {
	fp *os.File
}

func NewDeckWriter(fp *os.File) *DeckWriter {
	return &DeckWriter{fp: fp}
}

func (dw *DeckWriter) WriteDeck(d *Deck) error {
	err := dw.writeToCSV(d)
	if err != nil {
		return err
	}

	return nil
}

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
