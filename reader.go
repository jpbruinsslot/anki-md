package main

import "io"

// reader is a buffered rune reader, which allows us to unread runes for more
// than one position, in this case 3.
type reader struct {
	r   io.RuneScanner
	i   int // the position of the buffer
	n   int // how many runes are on the buffer
	pos Pos // last rune position, helps us with relaying location of syntax errors
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

// read will read the next rune from the io.RuneScanner
func (r *reader) read() (ch rune, pos Pos) {
	// When there are unread characters then read them of buffer first
	if r.n > 0 {
		r.n--
		return r.curr()
	}

	// Read next rune from the RuneScanner (`\r' is carriage return)
	ch, _, err := r.r.ReadRune()
	if err != nil {
		ch = eof
	} else if ch == '\r' {
		if ch, _, err := r.r.ReadRune(); err != nil {
			// nop
		} else if ch != '\n' {
			_ = r.r.UnreadRune()
		}
		ch = '\n'
	}

	// Update index, character, and position of the buffer.
	//
	// Using modulo to use beginning of the buffer (which is a slice of length
	// 3), so when we've used up all the positions of the buffer it'll start
	// at position 0.
	r.i = (r.i + 1) % len(r.buf)
	buf := &r.buf[r.i]
	buf.ch, buf.pos = ch, r.pos

	// Update Pos (position) of the buffer, increase line
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

// unreadRepeat is a wrapper method for `unread` to execute it `n` times.
func (r *reader) unreadRepeat(n int) {
	for i := 1; i <= n; i++ {
		r.unread()
	}
}
