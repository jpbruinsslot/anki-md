// Sources:
// - https://blog.gopheracademy.com/advent-2014/parsers-lexers/
// - https://github.com/benbjohnson/sql-parser
package anki

import (
	"bufio"
	"bytes"
	"io"
)

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

	if isPercent(ch) {
		// s.unread()
		return s.scanField()
	}

	// else if isHyphen(ch) {
	// 	// s.unread()
	// 	return s.scanCard()
	// }

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
		s.unread()
	}
}

func (s *Scanner) scanField() (token Token, lit string) {
	// Create buffer and read the current character into it
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// The next character should be another percent
	ch := s.read()
	if !isPercent(ch) {
		s.unread()
		return ILLEGAL, string(ch)
	}

	// Read until the next double percent
	for {
		ch = s.read()
		if ch == eof {
			break
		} else if isPercent(ch) {
			if ch = s.read(); isPercent(ch) {
				s.unreadRepeat(2)
				break
			}
			s.unread()
		}

		_, _ = buf.WriteRune(ch)
	}

	return FIELD, string(buf.String())

}

// func (s *Scanner) scanCard() (token Token, lit string) {
// 	// Create bugger and read the current character into it
// }

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
