package models

type Text struct {
	Text         string
	Axis         Axis
	TextAligment TextAligment
}

func NewText(text string, axis Axis, aligment TextAligment) *Text {
	return &Text{Text: text, Axis: axis, TextAligment: aligment}
}
