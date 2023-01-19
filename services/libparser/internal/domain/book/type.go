package book

type Book struct {
	Title   string
	Path    string
	Tags    []string
	Authors []string
	Size    uint64
}

func New(title string, path string, size uint64, tags []string, authors []string) Book {
	return Book{
		Title:   title,
		Path:    path,
		Tags:    tags,
		Authors: authors,
		Size:    size,
	}
}
