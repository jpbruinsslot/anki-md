package main

import (
	"bufio"
	"bytes"
	"io"
)

// eof represents a marker rune for the end of the reader, defining it here
// gives us a possibility to recognize it in the Scanner
var eof = rune(0)

type Scanner struct {
	r *reader
}

// NewScanner creates a Scanner struct, which contains a buffered rune reader.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: &reader{r: bufio.NewReader(r)}}
}

// Scan will scan individual runes, and identify if a specific rune is ...
// encountered
func (s *Scanner) Scan() (token Token, pos Pos, lit string) {

	// Read the next rune
	ch, pos := s.r.read()

	switch {
	case isWhitespace(ch):
		return s.scanWhitespace()
	case isPercent(ch):
		return s.scanField()
	case isHyphen(ch):
		s.r.unread()
		return s.scanCard()
	default:
		// Otherwise read the individual character
		switch ch {
		case eof:
			return EOF, pos, ""
		}
	}

	return ILLEGAL, pos, string(ch)
}

// scanField will scan the FIELD Token, as well as return the literal
// string contained in that FIELD Token.
func (s *Scanner) scanField() (token Token, pos Pos, lit string) {
	// Save the position of the field
	_, pos = s.r.curr()

	// Create buffer, here we'll write the runes into
	var buf bytes.Buffer

	// The next character should be another percent
	ch, _ := s.r.read()
	if !isPercent(ch) {
		s.r.unread()
		return ILLEGAL, pos, ""
	}

	// We read until we see the first non-whitespace character
	for {
		if ch, _ = s.r.read(); !isWhitespace(ch) {
			s.r.unread()
			break
		}
	}

	// Read until:
	// * double percent
	// * eof
	// * card
	for {
		ch, _ = s.r.read()
		if ch == eof {
			break
		} else if isPercent(ch) {

			// Peak ahead
			chNext, _ := s.r.read()
			s.r.unread()

			if isPercent(chNext) {
				s.r.unread()
				break
			}
		} else if isHyphen(ch) {

			// We start at depth 1, because
			// we already read the first hyphen
			if s.isCard(1) {
				break
			}

			// isCard will step back one too much, we
			// want to write the initial hyphen
			s.r.read()
		}

		// Write runes into buffer
		_, _ = buf.WriteRune(ch)
	}

	return FIELD, pos, buf.String()
}

// scanCard will scan the CARD Token
func (s *Scanner) scanCard() (token Token, pos Pos, lit string) {
	// Save the position of the field
	_, pos = s.r.curr()

	// Create buffer, here we'll write the runes into
	var buf bytes.Buffer

	// When a hyphen is found it should be three consecutive hyphens
	if !s.isCard(0) {
		return ILLEGAL, pos, ""
	}

	// Read until:
	// * eof
	// * not a hyphen
	// * more than 3 consecutive hyphens
	i := 0
	for {
		ch, _ := s.r.read()
		if ch == eof {
			break
		} else if !isHyphen(ch) {
			s.r.unread()
			break
		} else if i > 2 {
			s.r.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
			i++
		}
	}

	return CARD, pos, buf.String()
}

// scanWhiteSpace will WHITESPACE Tokens
func (s *Scanner) scanWhitespace() (token Token, pos Pos, lit string) {
	// Save the position of the field
	_, pos = s.r.curr()

	// Create buffer, here we'll write the runes into
	var buf bytes.Buffer

	// Read until:
	// * eof
	// * discontinuation of whitespaces
	for {
		ch, _ := s.r.read()
		if ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.r.unread()
			break
		}

		_, _ = buf.WriteRune(ch)
	}

	return WS, pos, buf.String()
}

// isCard will identify if a encountered `-` (hyphen) is part of a CARD Token,
// the next three characters should also be `-`
func (s *Scanner) isCard(depth int) bool {
	ch, _ := s.r.read()

	if !isHyphen(ch) || depth > 2 {
		s.r.unreadRepeat(depth + 1)
		return false
	}

	if depth < 2 {
		return s.isCard(depth + 1)
	}

	s.r.unreadRepeat(depth + 1)
	return true
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isPercent(ch rune) bool {
	return ch == '%'
}

func isHyphen(ch rune) bool {
	return ch == '-'
}
