package persistence

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/thayen/sample-rest/domain/model"
	"log"
)

const (
	dialect = "postgres"
)

type ContactRepository interface {
	FindAll() ([]Contact, error)
	Get(id int) (Contact, error)
	Create(contact Contact) (int, error)
	Update(contact Contact) error
	Delete(id int) error
	FindByEmail(email string) (Contact, error)
}
type contactSqlRepository struct {
	db *gorm.DB
}

func NewContactSqlRepository(conn string) (ContactRepository, error) {
	db, err := gorm.Open(dialect, conn)
	if err != nil {
		log.Println("Cannot get gorm connection", err)
		return nil, err
	}

	db.AutoMigrate(Contact{})
	return &contactSqlRepository{db: db}, nil
}

func (repo *contactSqlRepository) FindAll() ([]Contact, error) {
	contacts := make([]Contact, 0, 10)
	err := repo.db.Find(&contacts).Error
	return contacts, err
}

func (repo *contactSqlRepository) FindByEmail(email string) (Contact, error) {
	ct := Contact{}
	err := repo.db.Where("email = ? ", email).First(&ct).Error
	return ct, err
}

func (repo *contactSqlRepository) Get(id int) (Contact, error) {
	ct := Contact{}
	first := repo.db.Where("contact_id = ?", id).First(&ct)
	if first.RecordNotFound() {
		return ct, model.NotFound{}
	}
	err := first.Error
	return ct, err
}

func (repo *contactSqlRepository) Create(contact Contact) (int, error) {
	err := repo.db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&contact).Error
	})
	return contact.ID, err

}

func (repo *contactSqlRepository) Update(contact Contact) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		db := tx.First(&contact, "contact_id = ?", contact.ID)

		if db.RecordNotFound() {
			return model.NotFound{}
		}
		return db.Update(&contact).Error
	})
}

func (repo *contactSqlRepository) Delete(id int) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		db := tx.Delete(Contact{}, "contact_id = ? ", id)
		if db.RecordNotFound() {
			return model.NotFound{}
		}
		return db.Error
	})
}
