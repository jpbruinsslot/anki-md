package anki

import (
	"bufio"
	"bytes"
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
			if isCard(s, 1) {
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

func (s *Scanner) scanCard() (token Token, pos Pos, lit string) {
	// Save the position of the field
	_, pos = s.r.curr()

	// Create buffer, here we'll write the runes into
	var buf bytes.Buffer

	// When a hyphen is found it should be three consecutive hyphens
	if !isCard(s, 0) {
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

// FIXME: make it a struct method
// Given a `-` (hyphen), the next three characters should also be `-`
func isCard(s *Scanner, depth int) bool {
	ch, _ := s.r.read()

	// fmt.Printf("d: %d, ch: %s\n", depth, string(ch))

	if !isHyphen(ch) || depth > 2 {
		s.r.unreadRepeat(depth + 1)
		return false
	}

	if depth < 2 {
		return isCard(s, depth+1)
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

	// Update index, character, and position of the buffer
	r.i = (r.i + 1) % len(r.buf) // ???
	buf := &r.buf[r.i]
	buf.ch, buf.pos = ch, r.pos

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

func (r *reader) unreadRepeat(times int) {
	for i := 1; i <= times; i++ {
		r.unread()
	}
}
