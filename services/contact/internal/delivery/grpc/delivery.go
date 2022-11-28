package grpc

import (
	contact "smartdom/services/contact/internal/delivery/grpc/interface"
	"smartdom/services/contact/internal/useCase"
)

type Delivery struct {
	contact.UnimplementedContactServiceServer
	ucContact useCase.Contact
	ucGroup   useCase.Group

	options Options
}

type Options struct{}

func New(ucContact useCase.Contact, ucGroup useCase.Group, o Options) *Delivery {
	var d = &Delivery{
		ucContact: ucContact,
		ucGroup:   ucGroup,
	}

	d.SetOptions(o)
	return d
}

func (d *Delivery) SetOptions(options Options) {
	if d.options != options {
		d.options = options
	}
}
