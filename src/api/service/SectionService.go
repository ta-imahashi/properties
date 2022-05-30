package service

import (
	"strconv"

	"github.com/ta-imahashi/properties/model"
	"github.com/ta-imahashi/properties/repository"
)

type SectionService struct {
	SectionRepository repository.SectionRepository
	CounterRepository repository.CounterRepository
}

func NewSectionService(sectionRe repository.SectionRepository, counterRe repository.CounterRepository) SectionService {
	return SectionService{SectionRepository: sectionRe, CounterRepository: counterRe}
}

func (se *SectionService) List() []model.Section {
	return se.SectionRepository.Scan()
}

func (se *SectionService) Find(id string) model.Section {
	return se.SectionRepository.GetItem(id)
}

func (se *SectionService) GetSections(id string) []model.Section {
	return se.SectionRepository.Query(id)
}

func (se *SectionService) Create(section model.Section) model.Section {
	counter := se.CounterRepository.Increment(se.SectionRepository.TableName)
	section.Id, _ = strconv.Atoi(counter.Count)
	return se.SectionRepository.PutItem(section)
}

func (se *SectionService) Update(section model.Section) model.Section {
	return se.SectionRepository.UpdateItem(section)
}

func (se *SectionService) Delete(id string) {
	se.SectionRepository.DeleteItem(id)
}
