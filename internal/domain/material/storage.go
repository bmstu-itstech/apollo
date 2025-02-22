package material

import "errors"

// Storage is a generic interface for interaction with application data
type Storage interface {
	Materials(disciplineId int) ([]Material, error)
	Material(uuid string) (Material, error)
	Upsert(uuid string, material Material) error
	Disciplines() ([]Discipline, error)
	Discipline(id int) (Discipline, error)
	Departments() ([]Department, error)
	Department(id int) (Department, error)
}

var ErrMatNotExist = errors.New("material doesn't exist")
var ErrDeptNotExist = errors.New("department doesn't exist")
var ErrDiscNotExist = errors.New("discipline doesn't exist")
