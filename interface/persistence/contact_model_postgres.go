package persistence

import (
	"time"
)

type Contact struct {
	ID        int       `gorm:"AUTO_INCREMENT;column:contact_id"`
	Name      string    `gorm:"type:text;not_null"`
	Email     string    `gorm:"type:text;not_null"`
	BirthDate time.Time `gorm:"type:TIMESTAMP;not null;"`
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
	return c.ID
}
