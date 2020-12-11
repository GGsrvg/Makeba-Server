package models

type DiagramItem struct {
	Name        string
	Color       Color
	NumberVotes uint
}

func NewDiagramItem(name string, color Color, numberVotes uint) *DiagramItem {
	return &DiagramItem{name, color, numberVotes}
}
