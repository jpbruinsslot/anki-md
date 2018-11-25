package main_test

import (
	"reflect"
	"strings"
	"testing"

	anki "github.com/erroneousboat/anki-md"
)

func TestParser_ParseDeck(t *testing.T) {
	var tests = []struct {
		s   string
		d   *anki.Deck
		err string
	}{
		{
			s: `
%% This is a field 001
%% This is a field 002
---
%% This is a field 003
%% This is a field 004
---`,
			d: &anki.Deck{
				Cards: []anki.Card{
					anki.Card{
						Fields: []anki.Field{
							anki.Field{
								Content: "This is a field 001\n",
							},
							anki.Field{
								Content: "This is a field 002\n",
							},
						},
					},
					anki.Card{
						Fields: []anki.Field{
							anki.Field{
								Content: "This is a field 003\n",
							},
							anki.Field{
								Content: "This is a field 004\n",
							},
						},
					},
				},
			},
		},
	}

	for i, tt := range tests {
		d, err := anki.NewParser(strings.NewReader(tt.s)).ParseDeck()
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.d, d) {
			t.Errorf("%d. %q\n\ndeck mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.d, d)
		}
	}

}

func TestParser_ParseDeck_CardCount(t *testing.T) {
	s := `
	%% This is a field

	%% This is a field

	---

	%% This is a field

	%% This is a field

	---`

	d, err := anki.NewParser(strings.NewReader(s)).ParseDeck()
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	} else if len(d.Cards) != 2 {
		t.Fatalf("unexpected card count: %d", len(d.Cards))
	}
}

func TestParser_ParseDeck_CardCountTermination(t *testing.T) {
	s := `
	%% This is a field

	%% This is a field

	---

	%% This is a field

	%% This is a field`

	d, err := anki.NewParser(strings.NewReader(s)).ParseDeck()
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	} else if len(d.Cards) != 2 {
		t.Fatalf("unexpected card count: %d", len(d.Cards))
	}
}

func TestParser_ParseDeck_FieldCount(t *testing.T) {
	s := `
	%% This is a field

	%% This is a field

	---`

	d, err := anki.NewParser(strings.NewReader(s)).ParseDeck()
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	} else if len(d.Cards[0].Fields) != 2 {
		t.Fatalf("unexpected field count: %d", len(d.Cards[0].Fields))
	}
}

func TestParser_ParseDeck_Empty(t *testing.T) {
	s := ``

	d, err := anki.NewParser(strings.NewReader(s)).ParseDeck()
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	} else if len(d.Cards) != 0 {
		t.Fatalf("unexpected card count: %d", len(d.Cards))
	}
}

// errstring returns the string representation of an error.
func errstring(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
