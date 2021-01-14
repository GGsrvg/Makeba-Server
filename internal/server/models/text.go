package models

type Text struct {
	Value string

	/*
		FUTURE:
			need add link
			for create partions

		Example:
			- Harry Potter		tag = "harry_potter"	link = "https://example.com:8080/harry_potter/"
			- - First Book		tag = "first_book"		link = "https://example.com:8080/harry_potter/first_book/"
			- - Second Book		tag = "second_book"		link = "https://example.com:8080/harry_potter/second_book/"
			- - Third Book		tag = "third_book"		link = "https://example.com:8080/harry_potter/third_book/"
	*/
	// Tag  *string

	// Axis         Axis
	// TextAligment TextAligment
}

func NewText(text string) *Text {
	return &Text{Value: text}
}
