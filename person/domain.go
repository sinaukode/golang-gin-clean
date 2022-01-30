package person

import (
	"go-clean-code-gin/entity"
)

type PersonRepository interface {
	Create(person *entity.Person) (*entity.Person, error)
	GetAll() (*[]entity.Person, error)
	GetById(id int) (*entity.Person, error)
	Update(id int, person *entity.Person) (*entity.Person, error)
	Delete(id int) error
}

type PersonService interface {
	Create(person *entity.Person) (*entity.Person, error)
	GetAll() (*[]entity.Person, error)
	GetById(id int) (*entity.Person, error)
	Update(id int, person *entity.Person) (*entity.Person, error)
	Delete(id int) error
}
