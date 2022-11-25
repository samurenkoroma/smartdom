package postgres

import (
	"github.com/google/uuid"

	"smartdom/pkg/type/queryParameter"
	"smartdom/services/contact/internal/domain/contact"
)

func (r *Repository) CreateContact(contacts ...*contact.Contact) ([]*contact.Contact, error) {
	panic("implement me")
}

func (r *Repository) UpdateContact(ID, updateFn func(c *contact.Contact) (*contact.Contact, error)) (*contact.Contact, error) {
	panic("implement me")
}

func (r *Repository) DeleteContact(ID uuid.UUID) error {
	panic("implement me")
}

func (r *Repository) ListContact(parameter queryParameter.QueryParameter) ([]*contact.Contact, error) {
	panic("implement me")
}

func (r *Repository) ReadContactByID(ID uuid.UUID) (response *contact.Contact, err error) {
	panic("implement me")
}

func (r *Repository) CountContact() (uint64, error) {
	panic("implement me")
}
