package material

import "errors"

type Discipline struct {
	Id   int
	Name string // e.g. "Основы Информатики"
}

var ErrDiscNotExist = errors.New("discipline doesn't exist")
