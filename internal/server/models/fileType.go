package models

type FileType string

const (
	BOOK  FileType = "BOOK"  // pdf
	TABLE FileType = "TABLE" // excel
	IMAGE FileType = "IMAGE" // png, jpg, jpeg
)
