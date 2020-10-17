package models

type Text struct {
	Value string
	// Axis         Axis
	// TextAligment TextAligment
}

func NewText(text string) *Text {
	return &Text{Value: text}
}
