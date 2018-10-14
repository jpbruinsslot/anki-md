package anki_test

import (
	"strings"
	"testing"

	"github.com/erroneousboat/anki-md"
)

func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s   string
		tok anki.Token
		lit string
	}{
		{s: ``, tok: anki.EOF},
		{s: `%%`, tok: anki.FIELD},
	}

	for i, tt := range tests {
		s := anki.NewScanner(strings.NewReader(tt.s))
		tok, lit := s.Scan()

		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)
		} else if tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
		}
	}
}
