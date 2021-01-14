package models

import "time"

/*
	Type `Doucument` need for work with images, pdf, words and another
	if Document.Type == IMAGE need preview document
*/
type Document struct {
	Title       string
	Description string
	DateCreated time.Time
	DateUpdated time.Time
	Link        string
	Type        FileType
}
