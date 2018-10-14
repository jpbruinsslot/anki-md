package anki

type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	WS

	FIELD
	CARD
)
