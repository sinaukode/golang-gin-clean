package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-clean-code-gin/entity"
	"go-clean-code-gin/person"
)

type PersonRepositoryImpl struct {
	DB *gorm.DB
}

func CreatePersonRepository(DB *gorm.DB) person.PersonRepository {
	return &PersonRepositoryImpl{DB}
}

func (p *PersonRepositoryImpl) Create(person *entity.Person) (*entity.Person, error) {
	err := p.DB.Save(&person).Error
	if err != nil {
		fmt.Printf("[PersonRepoImpl.Create] error execute query %v \n", err)
		return nil, err
	}
	return person, nil
}

func (p *PersonRepositoryImpl) GetAll() (*[]entity.Person, error) {
	var persons []entity.Person
	err := p.DB.Find(&persons).Error
	if err != nil {
		fmt.Printf("[PersonRepoImpl.GetAll] error execute query %v \n", err)
		return nil, err
	}
	return &persons, nil
}

func (p *PersonRepositoryImpl) GetById(id int) (*entity.Person, error) {
	var person entity.Person
	err := p.DB.Where("id = ?", id).First(&person).Error

	if err != nil {
		fmt.Printf("[PersonRepoImpl.GetById] error execute query %v \n", err)
		return nil, err
	}
	return &person, nil
}

func (p *PersonRepositoryImpl) Update(id int, person *entity.Person) (*entity.Person, error) {

	var UpdatePerson = entity.Person{}

	err := p.DB.Where("id = ?", id).First(&UpdatePerson).Update(&person).Error
	if err != nil {
		fmt.Printf("[PersonRepoImpl.Update] error execute query %v \n", err)
		return nil, err
	}
	return &UpdatePerson, nil
}

func (p *PersonRepositoryImpl) Delete(id int) error {
	var person = entity.Person{}
	err := p.DB.Where("id = ?", id).First(&person).Delete(&person).Error
	if err != nil {
		fmt.Printf("[PersonRepoImpl.Delete] error execute query %v \n", err)
		return err
	}

	return nil
}
