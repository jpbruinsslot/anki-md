// https://github.com/influxdata/influxql/blob/master/scanner.go
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
	r *reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: &reader{r: bufio.NewReader(r)}}
}

func (s *Scanner) Scan() (token Token, pos Pos, lit string) {

	// Read the next rune
	ch, pos := s.r.read()

	switch {
	case isPercent(ch):
		s.r.unread()
		return s.scanField()
	case isHyphen(ch):
		s.r.unread()
		return s.scanCard()
	default:
		// Otherwise read the individual character
		switch ch {
		case eof:
			return EOF, ""
		}
	}

	return ILLEGAL, string(ch)
}

func (s *Scanner) scanField() (token Token, lit string) {
	// Save the position of the field
	_, pos := s.r.read()
	s.r.unread()

	// Create buffer, here we'll write the runes into
	var buf bytes.Buffer

	// ch := s.read()
	// for isPercent(ch) {
	// 	ch = s.read()
	// }

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

// reader is a buffered rune reader
type reader struct {
	r   io.RuneScanner
	i   int // at what position of the reader is the buffer (correct?)
	n   int // how many character are on the buffer
	pos Pos // last rune position
	buf [3]struct {
		ch  rune
		pos Pos
	}
	eof bool // true if reader has has ever seen eof
}

// ReadRune implements the io.RuneScanner interface, it reads the next
// rune from the reader. It doesn't return the size.
func (r *reader) ReadRune() (ch rune, size int, err error) {
	ch, _ = r.read()
	if ch == eof {
		err = io.EOF
	}
	return
}

// UnReadRune implements the io.RuneScanner interface, it unreads the
// previously read rune back onto the buffer
func (r *reader) UnReadRune() error {
	r.unread()
	return nil
}

func (r *reader) read() (ch rune, pos Pos) {
	// When there are unread characters then read them of buffer first
	if r.n > 0 {
		r.n--
		return r.curr()
	}

	// Read next rune from the RuneScanner
	ch, _, err := r.r.ReadRune()
	if err != nil {
		ch = eof
	} else if ch == '\r' { // what is /r?
		if ch, _, err := r.r.ReadRune(); err != nil {
			// nop
		} else if ch != '\n' {
			_ = r.r.UnreadRune()
		}
		ch = '\n'
	}

	// Update index of the buffer
	r.i = (r.i + 1) % len(r.buf) // ???

	// Update position of the buffer, increase line
	// when newline is encountered
	if ch == '\n' {
		r.pos.Line++
		r.pos.Char = 0
	} else if !r.eof {
		r.pos.Char++
	}

	// Mark the reader as EOF
	if ch == eof {
		r.eof = true
	}

	return r.curr()
}

// curr returns the last read character and position
func (r *reader) curr() (ch rune, pos Pos) {

	// Get character from buffer
	i := (r.i - r.n + len(r.buf)) % len(r.buf)
	buf := &r.buf[i]

	return buf.ch, buf.pos
}

// unread pushes the previously read rune back onto the buffer
func (r *reader) unread() {
	r.n++
}
