package anki_test

import (
	"strings"
	"testing"

	"github.com/erroneousboat/anki-md"
)

func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s   string     // The input string
		tok anki.Token // The token that should be returned
		lit string     // The literal string the token represents
	}{
		// Special tokens
		// {s: ``, tok: anki.EOF},
		//
		// {s: `%%`, tok: anki.FIELD},
		// {s: `%%This is a field`, tok: anki.FIELD, lit: "This is a field"},
		// {s: `%%This is a field%`, tok: anki.FIELD, lit: "This is a field%"},
		// {s: `%%This is a field% `, tok: anki.FIELD, lit: "This is a field% "},
		// {s: `%%This is a field%%`, tok: anki.FIELD, lit: "This is a field"},

		{s: `---`, tok: anki.CARD, lit: "---"},
		{s: `--x`, tok: anki.ILLEGAL},
		// {s: `-xx`, tok: anki.ILLEGAL},
		// {s: `----`}
		// {s: `--`}
		// {s: `-`}
	}

	for i, tt := range tests {
		s := anki.NewScanner(strings.NewReader(tt.s))
		tok, pos, lit := s.Scan()

		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: pos=%q exp=%q got=%q <%q>", i, tt.s, pos, tt.tok, tok, lit)
		} else if tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: pos=%q exp=%q got=%q", i, tt.s, pos, tt.lit, lit)
		}
	}
}
