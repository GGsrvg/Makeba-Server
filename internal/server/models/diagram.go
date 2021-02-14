package models

type Diagram struct {
	Type  DiargamType
	Items []DiagramItem
}

func NewDiagram(diagramType DiargamType, items ...DiagramItem) *Diagram {
	return &Diagram{
		Type:  diagramType,
		Items: items}
}
