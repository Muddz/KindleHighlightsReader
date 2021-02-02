package highlight

type Highlight struct {
	Title  string
	Author string
	Text   string
}

func New(title string, author string, text string) Highlight {
	return Highlight{
		Title:  title,
		Author: author,
		Text:   text,
	}
}
