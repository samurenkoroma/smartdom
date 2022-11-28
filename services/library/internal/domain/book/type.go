package book

import (
	"github.com/google/uuid"
	"smartdom/services/library/internal/domain/book/path"
	"smartdom/services/library/internal/domain/book/title"
)

type Book struct {
	id    uuid.UUID
	title title.Title
	path  path.Path
}

func (b Book) ID() uuid.UUID {
	return b.id
}

func (b Book) Title() title.Title {
	return b.title
}

func (b Book) Path() path.Path {
	return b.path
}
