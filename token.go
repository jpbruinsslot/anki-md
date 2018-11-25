package main

type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	WS

	// Keywords
	FIELD
	CARD
)

// Pos represents the line and character position of a token
type Pos struct {
	Line int
	Char int
}
