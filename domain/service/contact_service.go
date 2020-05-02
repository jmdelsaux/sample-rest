package service

import (
	"github.com/thayen/sample-rest/interface/persistence"
	"log"
)

//ContactService
type ContactService interface {
	//Duplicate returns true if the email is already knowed
	Duplicate(email string) bool
}

type defaultContactService struct {
	repo persistence.ContactRepository
}

//Create a new ContactService
func NewContactService(repository persistence.ContactRepository) ContactService {
	return &defaultContactService{
		repo: repository,
	}
}

func (u *defaultContactService) Duplicate(email string) bool {
	byEmail, err := u.repo.FindByEmail(email)
	if err != nil {
		log.Print(err)
		return false
	}
	return byEmail.Email != ""
}
