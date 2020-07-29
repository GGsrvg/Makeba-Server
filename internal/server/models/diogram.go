package models

type Diogram struct {
	Data []int
}

func NewDiogram(data ...int) *Diogram {
	return &Diogram{Data: data}
}
