// Both helper methods `scan` and `unscan` allows us to leverage the buffer,
// when for instance we would encounter a token that was not allowed in a
// particular sequence, or parsing error occurred, then it is necessary to
// `unscan` that token.
package main

import "io"

type Parser struct {
	s   *Scanner
	buf struct {
		tok Token
		lit string
		n   int
	}
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

func (p *Parser) ParseDeck() (*Deck, error) {
	cards := make([]Card, 0)
	fields := make([]Field, 0)

	for {
		// NOTE: use `scanIgnoreWhitespace()` instead of p.s.Scan()
		// when necessary.
		if tok, _, lit := p.s.Scan(); tok == EOF {

			// When there are still fields present, add them to the
			// card
			if len(fields) > 0 {
				cards = append(cards, Card{Fields: fields})
			}

			return &Deck{Cards: cards}, nil
		} else if tok == FIELD {
			fields = append(fields, Field{Content: lit})
		} else if tok == CARD {
			cards = append(cards, Card{Fields: fields})
			fields = make([]Field, 0)
		}
	}
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, _, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() {
	p.buf.n = 1
}

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}

	return
}
