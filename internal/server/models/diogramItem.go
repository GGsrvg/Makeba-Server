package models

type DiogramItem struct {
	Name        string
	Color       Color
	NumberVotes uint
}

func NewDiogramItem(name string, color Color, numberVotes uint) *DiogramItem {
	return &DiogramItem{name, color, numberVotes}
}
