package models

type Stat struct {
	Type  string
	Title string
	Model interface{}
}

func NewStatText(title string, model Text) *Stat {
	return &Stat{Type: "text", Title: title, Model: model}
}

func NewStatDiogram(title string, model Diogram) *Stat {
	return &Stat{Type: "diogram", Title: title, Model: model}
}
