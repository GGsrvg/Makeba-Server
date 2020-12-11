package models

type Diagram struct {
	Items []DiagramItem
}

func NewDiagram(items ...DiagramItem) *Diagram {
	return &Diagram{Items: items}
}
