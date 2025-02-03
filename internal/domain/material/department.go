package material

import "errors"

type Department struct {
	Id   int
	Name string // e.g. "ИУ-9", "ФН-12"
	Desc string // e.g. "Теоретическая информатика и компьютерные технологии"
}

// TODO: unify this and ErrDiscNotExist?
var ErrDeptNotExist = errors.New("department doesn't exist")
