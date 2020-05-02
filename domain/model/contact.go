package model

import "time"

//Contact Simple model interface
type Contact interface {
	//GetId return Id of the contact
	GetId() int
	//GetName return Name of the contact
	GetName() string
	//GetEmail return Email of the contact
	GetEmail() string
	//GetBirthDate return BirthDate of the contact
	GetBirthDate() time.Time
}
