package mock

import (
	"github.com/bmstu-itstech/apollo/internal/domain/material"
)

// MockStorage is an implementation of material.Storage using maps,
// intended for basic testing.
type MockStorage struct {
	materials   map[string]material.Material
	departments map[int]material.Department
	disciplines map[int]material.Discipline
}

func NewMockStorage() MockStorage {
	return MockStorage{
		make(map[string]material.Material),
		make(map[int]material.Department),
		make(map[int]material.Discipline),
	}
}

func (s *MockStorage) Materials(discipline material.Discipline) ([]material.Material, error) {
	materials := make([]material.Material, 0, len(s.materials))
	for _, m := range s.materials {
		if m.DisciplineId == discipline.Id {
			materials = append(materials, m)
		}
	}
	return materials, nil
}

func (s *MockStorage) Material(uuid string) (material.Material, error) {
	m, ok := s.materials[uuid]
	if !ok {
		return material.Material{}, material.ErrMatNotExist
	}
	return m, nil
}

func (s *MockStorage) Upsert(uuid string, material material.Material) error {
	s.materials[uuid] = material
	return nil
}

func (s *MockStorage) Disciplines() ([]material.Discipline, error) {
	disciplines := make([]material.Discipline, 0, len(s.disciplines))
	for _, d := range s.disciplines {
		disciplines = append(disciplines, d)
	}
	return disciplines, nil
}

func (s *MockStorage) Discipline(id int) (material.Discipline, error) {
	d, ok := s.disciplines[id]
	if !ok {
		return material.Discipline{}, material.ErrDiscNotExist
	}
	return d, nil
}

func (s *MockStorage) Departments() ([]material.Department, error) {
	departments := make([]material.Department, 0, len(s.departments))
	for _, d := range s.departments {
		departments = append(departments, d)
	}
	return departments, nil
}

func (s *MockStorage) Department(id int) (material.Department, error) {
	d, ok := s.departments[id]
	if !ok {
		return material.Department{}, material.ErrDeptNotExist
	}
	return d, nil
}

func (s *MockStorage) AddDepartment(department material.Department) {
	s.departments[department.Id] = department
}
func (s *MockStorage) AddDiscipline(discipline material.Discipline) {
	s.disciplines[discipline.Id] = discipline
}
