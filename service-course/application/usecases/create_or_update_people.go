package usecases

import (
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/ports"
)

type CreateOrUpdatePeople struct {
	peopleRepository ports.PeopleRepository
}

func NewCreateOrUpdatePeople(PeopleRepository ports.PeopleRepository) *CreateOrUpdatePeople {
	return &CreateOrUpdatePeople{peopleRepository: PeopleRepository}
}

type CreatePeopleDTO struct {
	Id       string
	Name     string
	Type     string
	PhotoUrl string
}

func (cp *CreateOrUpdatePeople) Execute(Data CreatePeopleDTO) (*entities.People, error) {
	findPeople, err := cp.peopleRepository.FindById(Data.Id)
	if err != nil {
		return nil, err
	}

	var create = func() (*entities.People, error) {
		people, err := entities.NewPeople(Data.Id, Data.Name, Data.Type, nil)
		if err != nil {
			return nil, err
		}

		if err := cp.peopleRepository.Create(people); err != nil {
			return nil, err
		}

		return people, nil
	}

	var update = func() (*entities.People, error) {
		err := findPeople.Update(Data.Name, nil)
		if err != nil {
			return nil, err
		}

		if err := cp.peopleRepository.Update(findPeople); err != nil {
			return nil, err
		}

		return findPeople, nil
	}

	if findPeople != nil {
		return update()
	}

	return create()
}
