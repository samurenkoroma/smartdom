package contact

import (
	"github.com/google/uuid"
	"smartdom/pkg/type/queryParameter"
	"smartdom/services/contact/internal/domain/contact"
)

func (uc *UseCase) Create(contacts ...*contact.Contact) ([]*contact.Contact, error) {
	return uc.adapterStorage.CreateContact(contacts...)
}

func (uc *UseCase) Update(contactUpdate contact.Contact) (*contact.Contact, error) {
	return uc.adapterStorage.UpdateContact(contactUpdate.ID(), func(c *contact.Contact) (*contact.Contact, error) {
		return contact.NewWithID(
			contactUpdate.ID(),
			contactUpdate.CreatedAt(),
			contactUpdate.ModifiedAt(),
			contactUpdate.PhoneNumber(),
			contactUpdate.Email(),
			contactUpdate.Name(),
			contactUpdate.Surname(),
			contactUpdate.Patronymic(),
			contactUpdate.Age(),
			contactUpdate.Gender())
	})
}

func (uc *UseCase) Delete(ID uuid.UUID) error {
	return uc.adapterStorage.DeleteContact(ID)
}

func (uc *UseCase) List(parameter queryParameter.QueryParameter) ([]*contact.Contact, error) {
	return uc.adapterStorage.ListContact(parameter)
}

func (uc *UseCase) ReadByID(ID uuid.UUID) (response *contact.Contact, err error) {
	return uc.adapterStorage.ReadContactByID(ID)
}

func (uc *UseCase) Count() (uint64, error) {
	return uc.adapterStorage.CountContact()
}
