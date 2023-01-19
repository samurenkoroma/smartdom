package cli

import "smartdom/services/libparser/internal/useCase"

type Delivery struct {
	parser  useCase.LibraryParser
	options Options
}

type Options struct {
	Dir string
}

func New(parser useCase.LibraryParser, options Options) *Delivery {
	return &Delivery{
		parser:  parser,
		options: options,
	}
}

func (d Delivery) Parse() {
	books, err := d.parser.Parse(d.options.Dir)
	if err != nil {
		return
	}
	err = d.parser.Save(books)
	if err != nil {
		return
	}

}
