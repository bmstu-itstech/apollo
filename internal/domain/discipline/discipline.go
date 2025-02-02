package discipline

import "errors"

type Discipline struct {
	Id   int    `json:"id"`
	Name string `json:"name"` // e.g. "Основы Информатики"
}

var ErrNotExist = errors.New("discipline doesn't exist")
