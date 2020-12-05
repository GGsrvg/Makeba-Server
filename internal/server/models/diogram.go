package models

type Diogram struct {
	Items []DiogramItem
}

func NewDiogram(items ...DiogramItem) *Diogram {
	return &Diogram{Items: items}
}
