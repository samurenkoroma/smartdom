package slug

import "github.com/gosimple/slug"

type Slug string

func (s Slug) String() string {
	return string(s)
}

func New(value string) Slug {
	return Slug(slug.Make(value))
}
