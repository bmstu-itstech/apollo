package material

import "errors"

type Department struct {
	Id   int
	Name string // e.g. "ИУ-9", "ФН-12"
	Desc string // e.g. "Теоретическая информатика и компьютерные технологии"
}

func NewDepartment(id int, name, description string) (Department, error) {
	if id < 0 {
		return Department{}, errors.New("expected non-negative id")
	} else if name == "" {
		return Department{}, errors.New("expected non-empty name")
	} else if description == "" {
		return Department{}, errors.New("expected non-empty description")
	}
	return Department{Id: id, Name: name, Desc: description}, nil
}

func MustNewDepartment(id int, name, description string) Department {
	d, err := NewDepartment(id, name, description)
	if err != nil {
		panic(err)
	}
	return d
}
