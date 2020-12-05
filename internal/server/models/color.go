package models

type Color struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

func NewColor(r uint8, g uint8, b uint8) *Color {
	return &Color{Red: r, Green: g, Blue: b}
}
