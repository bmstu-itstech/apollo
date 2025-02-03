package material

import "errors"

// Storage contains all application data and provides methods for interacting with it
type Storage interface {
	Materials(discipline Discipline) ([]Material, error)
	Material(uuid string) (Material, error)
	Upsert(uuid string, material Material) error
	Disciplines() ([]Discipline, error)
	Discipline(id int) (Discipline, error)
	Departments() ([]Department, error)
	Department(id int) (Department, error)
}

// TODO: unify these errors?
var ErrDeptNotExist = errors.New("department doesn't exist")
var ErrDiscNotExist = errors.New("discipline doesn't exist")
