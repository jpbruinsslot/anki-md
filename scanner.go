package anki

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

// eof represents a marker rune for the end of the reader
var eof = rune(0)

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) Scan() (token Token, lit string) {

	// Read the next rune
	ch := s.read()

	// If we see a percent then ...
	// If we see a hyphen then ...
	if isPercent(ch) {
		return s.scanField()
	} else if isHyphen(ch) {
		s.unread()
		return s.scanCard()
	}

	// Otherwise read the individual character
	switch ch {
	case eof:
		return EOF, ""
	}

	return ILLEGAL, string(ch)
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}

	return ch
}

// Place the previously read rune back on the reader
func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
}

func (s *Scanner) unreadRepeat(times int) {
	for i := 1; i <= times; i++ {
		fmt.Println("unread")
		s.unread()
	}
}

func (s *Scanner) scanField() (token Token, lit string) {
	// Create buffer, here we'll write the runes into
	var buf bytes.Buffer

	// The next character should be another percent
	ch := s.read()
	if !isPercent(ch) {
		s.unreadRepeat(2)
		return ILLEGAL, string(ch)
	}

	// Read until the next double percent, or eof.
	// When it is found, place the reader two steps back
	for {
		ch = s.read()
		if ch == eof {
			break
		} else if isPercent(ch) {
			chNext := s.read()
			if isPercent(chNext) {
				s.unreadRepeat(2)
				break
			}
			s.unread()
		}

		// Write runes into buffer
		_, _ = buf.WriteRune(ch)
	}

	return FIELD, buf.String()
}

func (s *Scanner) scanCard() (token Token, lit string) {
	// Create buffer, here we'll write the runes into
	var buf bytes.Buffer

	// When a hyphen is found it should be three consecutive hyphens
	if !isCard(s, 0) {
		fmt.Println("!card")
		return ILLEGAL, string(s.read())
	}

	for {
		ch := s.read()
		fmt.Println(string(ch))
		if ch == eof {
			fmt.Println("eof")
			break
		} else if !isHyphen(ch) {
			fmt.Println("!hyphen")
			s.unread()
			break
		} else {
			fmt.Println("writeRune")
			_, _ = buf.WriteRune(ch)
		}
	}

	return CARD, buf.String()
}

// Given a found `-` hyphen, the next three characters should also be `-`
func isCard(s *Scanner, depth int) bool {
	ch := s.read()
	fmt.Printf("d: %d, ch: %s\n", depth, string(ch))

	if !isHyphen(ch) || depth > 2 {
		// TODO: return regular text
		// s.unreadRepeat(depth + 1)
		return false
	}

	if depth < 2 {
		return isCard(s, depth+1)
	}

	fmt.Println("true")
	// s.unreadRepeat(depth + 1)
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
