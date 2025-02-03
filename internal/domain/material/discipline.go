package material

import "errors"

type Discipline struct {
	Id   int
	Name string // e.g. "Основы Информатики"
}

func NewDiscipline(id int, name string) (Discipline, error) {
	if id < 0 {
		return Discipline{}, errors.New("expected non-negative id")
	} else if name == "" {
		return Discipline{}, errors.New("expected non-empty name")
	}
	return Discipline{Id: id, Name: name}, nil
}

func MustNewDiscipline(id int, name string) Discipline {
	d, err := NewDiscipline(id, name)
	if err != nil {
		panic(err)
	}
	return d
}
