package service

import (
	"go-clean-code-gin/entity"
	"go-clean-code-gin/person"
)

type PersonServiceImpl struct {
	personRepository person.PersonRepository
}

func CreatePersonService(personRepository person.PersonRepository) person.PersonService {
	return &PersonServiceImpl{personRepository}
}

func (p *PersonServiceImpl) Create(person *entity.Person) (*entity.Person, error) {
	return p.personRepository.Create(person)
}

func (p *PersonServiceImpl) GetAll() (*[]entity.Person, error) {
	return p.personRepository.GetAll()
}

func (p *PersonServiceImpl) GetById(id int) (*entity.Person, error) {
	return p.personRepository.GetById(id)
}

func (p *PersonServiceImpl) Update(id int, person *entity.Person) (*entity.Person, error) {
	return p.personRepository.Update(id, person)
}

func (p *PersonServiceImpl) Delete(id int) error {
	return p.personRepository.Delete(id)
}
