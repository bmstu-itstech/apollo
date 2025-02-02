package discipline

import "errors"

type Discipline struct {
	Id   int
	Name string // e.g. "Основы Информатики"
}

var ErrNotExist = errors.New("discipline doesn't exist")
