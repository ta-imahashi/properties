package service

import (
	"strconv"

	"github.com/ta-imahashi/properties/model"
	"github.com/ta-imahashi/properties/repository"
)

type PropertyService struct {
	PropertyRepository repository.PropertyRepository
	CounterRepository  repository.CounterRepository
}

func NewPropertyService(propertyRe repository.PropertyRepository, counterRe repository.CounterRepository) PropertyService {
	return PropertyService{PropertyRepository: propertyRe, CounterRepository: counterRe}
}

func (se *PropertyService) List() []model.Property {
	return se.PropertyRepository.Scan()
}

func (se *PropertyService) Find(id string) model.Property {
	return se.PropertyRepository.GetItem(id)
}

func (se *PropertyService) Create(property model.Property) model.Property {
	counter := se.CounterRepository.Increment(se.PropertyRepository.TableName)
	property.Id, _ = strconv.Atoi(counter.Count)
	return se.PropertyRepository.PutItem(property)
}

func (se *PropertyService) Update(property model.Property) model.Property {
	return se.PropertyRepository.UpdateItem(property)
}

func (se *PropertyService) Delete(id string) {
	se.PropertyRepository.DeleteItem(id)
}
