package http

import (
	"github.com/thayen/sample-rest/domain/model"
	"time"
)

type ContactCreation struct {
	Contact
}

type ContactUpdate struct {
	Contact
	Id int `json:"id"`
}

type Contact struct {
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birth_date"`
	Email     string    `json:"email"`
}

type ContactResponse struct {
	Contact
	Id int `json:"id"`
}

func toContactResponse(c model.Contact) ContactResponse {
	return ContactResponse{
		Contact: Contact{
			Name:      c.GetName(),
			BirthDate: c.GetBirthDate(),
			Email:     c.GetEmail(),
		},
		Id: c.GetId(),
	}
}

func (c ContactUpdate) GetId() int {
	return c.Id
}

func (c Contact) GetName() string {
	return c.Name
}

func (c Contact) GetEmail() string {
	return c.Email
}

func (c Contact) GetBirthDate() time.Time {
	return c.BirthDate
}

func (c Contact) GetId() int {
	panic("not implementef")
}
