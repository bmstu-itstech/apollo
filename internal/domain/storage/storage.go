package storage

import (
	"github.com/bmstu-itstech/apollo/internal/domain/department"
	"github.com/bmstu-itstech/apollo/internal/domain/discipline"
	"github.com/bmstu-itstech/apollo/internal/domain/material"
)

type Storage interface {
	Materials(discipline discipline.Discipline) ([]material.Material, error)
	Material(uuid string) (material.Material, bool)
	Upsert(uuid string, material material.Material) error
	Disciplines(department department.Department) ([]discipline.Discipline, error)
	Discipline(id int) (discipline.Discipline, bool)
	Departments() ([]department.Department, error)
	Department(id int) (department.Department, bool)
}
