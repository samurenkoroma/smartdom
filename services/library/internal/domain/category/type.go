package category

import (
	"github.com/google/uuid"
	"smartdom/pkg/type/slug"
	"smartdom/services/library/internal/domain/category/name"
)

type Category struct {
	id   uuid.UUID
	name name.Name
	slug slug.Slug
}

func New(id uuid.UUID, name name.Name) (*Category, error) {
	if id == uuid.Nil {
		id = uuid.New()
	}
	return &Category{
		id:   id,
		name: name,
		slug: slug.New(name.String()),
	}, nil
}
