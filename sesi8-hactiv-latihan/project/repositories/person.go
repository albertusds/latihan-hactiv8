package repositories

import (
	"project-sesi8/models"

	"github.com/jinzhu/gorm"
)

type PersonRepo interface {
	CreatePerson(person *models.Person) error
	GetAllPersons() (*[]models.Person, error)
	GetPersonById(id int) (*models.Person, error)
	DeletePersonById(id int) (*models.Person, error)
	UpdatePersonById(person *models.Person) (*models.Person, error)
}

type personRepo struct {
	db *gorm.DB
}

func NewPersonRepo(db *gorm.DB) PersonRepo {
	return &personRepo{db}
}

func (p *personRepo) CreatePerson(person *models.Person) error {
	return p.db.Create(person).Error
}

func (p *personRepo) GetAllPersons() (*[]models.Person, error) {
	var person []models.Person
	err := p.db.Find(&person).Error
	return &person, err
}

func (p *personRepo) GetPersonById(id int) (*models.Person, error) {
	var person models.Person

	err := p.db.First(&person, "id=?", id).Error
	return &person, err
}

func (p *personRepo) DeletePersonById(id int) (*models.Person, error) {
	var person models.Person
	var personRs models.Person

	err := p.db.First(&personRs, "id=?", id).Error
	if err == nil {
		err = p.db.Delete(&person, "id=?", id).Error
		return &personRs, err
	}
	return nil, err
}

func (p *personRepo) UpdatePersonById(person *models.Person) (*models.Person, error) {
	var prsn models.Person

	err := p.db.Model(&prsn).Where("id=?", &person.Model.ID).Updates(models.Person{FirstName: person.FirstName, LastName: person.LastName}).Error
	if err != nil {
		return nil, err
	}
	return person, err
}
