package usecase

import (
	"github.com/thayen/sample-rest/domain/model"
	"github.com/thayen/sample-rest/domain/service"
	"github.com/thayen/sample-rest/interface/persistence"
)

type ContactUsecase interface {
	Create(c model.Contact) (int, error)
	Update(c model.Contact) error
	Delete(id int) error
	Get(id int) (model.Contact, error)
	FindAll() ([]model.Contact, error)
}

//NewContactUsecase Create a new contact usecase
func NewContactUsecase(contactService service.ContactService, contactRepository persistence.ContactRepository) ContactUsecase {
	return &defaultContactUsecase{
		contactService:    contactService,
		contactRepository: contactRepository,
	}
}

type defaultContactUsecase struct {
	contactService    service.ContactService
	contactRepository persistence.ContactRepository
}

func (d *defaultContactUsecase) FindAll() ([]model.Contact, error) {
	all, err := d.contactRepository.FindAll()
	ct := make([]model.Contact, 0, len(all))
	for _, c := range all {
		ct = append(ct, c)
	}
	return ct, err
}

func (d *defaultContactUsecase) Create(c model.Contact) (int, error) {
	if d.contactService.Duplicate(c.GetEmail()) {
		return 0, model.EmailAlreadyExists{}
	}
	return d.contactRepository.Create(toContact(c))
}

func (d *defaultContactUsecase) Update(c model.Contact) error {
	return d.contactRepository.Update(toContact(c))
}

func (d *defaultContactUsecase) Delete(id int) error {
	return d.contactRepository.Delete(id)
}

func (d *defaultContactUsecase) Get(id int) (model.Contact, error) {
	return d.contactRepository.Get(id)
}

func toContact(c model.Contact) persistence.Contact {
	return persistence.Contact{
		Name:      c.GetName(),
		Email:     c.GetEmail(),
		BirthDate: c.GetBirthDate(),
	}
}
