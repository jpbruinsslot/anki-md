package anki_test

import (
	"strings"
	"testing"

	anki "github.com/erroneousboat/anki-md"
)

func TestParser_ParseDeck(t *testing.T) {
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

func TestParser_ParseDeck_CardCount(t *testing.T) {
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
