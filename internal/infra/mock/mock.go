package mock

import (
	"sync"

	"github.com/bmstu-itstech/apollo/internal/domain/material"
)

// MockStorage is an implementation of material.Storage using maps,
// intended for basic testing.
type MockStorage struct {
	sync.RWMutex
	materials   map[string]material.Material
	departments map[int]material.Department
	disciplines map[int]material.Discipline
}

func NewMockStorage() *MockStorage {
	return &MockStorage{
		materials:   make(map[string]material.Material),
		departments: make(map[int]material.Department),
		disciplines: make(map[int]material.Discipline),
	}
}

func (s *MockStorage) Materials(disciplineId int) ([]material.Material, error) {
	s.RLock()
	defer s.RUnlock()
	materials := make([]material.Material, 0, len(s.materials))
	for _, m := range s.materials {
		if m.DisciplineId == disciplineId {
			materials = append(materials, m)
		}
	}
	return materials, nil
}

func (s *MockStorage) Material(uuid string) (material.Material, error) {
	s.RLock()
	defer s.RUnlock()
	m, ok := s.materials[uuid]
	if !ok {
		return material.Material{}, material.ErrMatNotExist
	}
	return m, nil
}

func (s *MockStorage) Upsert(uuid string, material material.Material) error {
	s.Lock()
	defer s.Unlock()
	m, ok := s.materials[uuid]
	if ok {
		material.Views = m.Views
		material.Created = m.Created
	}
	s.materials[uuid] = material
	return nil
}

func (s *MockStorage) Disciplines() ([]material.Discipline, error) {
	s.RLock()
	defer s.RUnlock()
	disciplines := make([]material.Discipline, 0, len(s.disciplines))
	for _, d := range s.disciplines {
		disciplines = append(disciplines, d)
	}
	return disciplines, nil
}

func (s *MockStorage) Discipline(id int) (material.Discipline, error) {
	s.RLock()
	defer s.RUnlock()
	d, ok := s.disciplines[id]
	if !ok {
		return material.Discipline{}, material.ErrDiscNotExist
	}
	return d, nil
}

func (s *MockStorage) Departments() ([]material.Department, error) {
	s.RLock()
	defer s.RUnlock()
	departments := make([]material.Department, 0, len(s.departments))
	for _, d := range s.departments {
		departments = append(departments, d)
	}
	return departments, nil
}

func (s *MockStorage) Department(id int) (material.Department, error) {
	s.RLock()
	defer s.RUnlock()
	d, ok := s.departments[id]
	if !ok {
		return material.Department{}, material.ErrDeptNotExist
	}
	return d, nil
}

func (s *MockStorage) AddDepartment(department material.Department) {
	s.Lock()
	defer s.Unlock()
	s.departments[department.Id] = department
}
func (s *MockStorage) AddDiscipline(discipline material.Discipline) {
	s.Lock()
	defer s.Unlock()
	s.disciplines[discipline.Id] = discipline
}
