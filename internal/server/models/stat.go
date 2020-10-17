package models

type Stat struct {
	Title string
	// the following attributes only one of them can have a value
	Text    *Text
	Diogram *Diogram
}

func NewStatText(title string, model *Text) *Stat {
	return &Stat{Title: title, Text: model}
}

func NewStatDiogram(title string, model *Diogram) *Stat {
	return &Stat{Title: title, Diogram: model}
}
