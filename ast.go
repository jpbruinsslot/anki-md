package anki

type Deck struct {
	Cards []Card
}

type Card struct {
	Fields []Field
}

type Field struct {
	Content string
}
