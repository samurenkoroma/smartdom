package tag

import (
	"smartdom/pkg/type/slug"
)

type Tag struct {
	name string
	slug slug.Slug
}

func (t Tag) Name() string {
	return string(t.name)
}

func New(value string) (*Tag, error) {
	t := Tag{
		name: value,
		slug: slug.New(value),
	}
	return &t, nil
}
