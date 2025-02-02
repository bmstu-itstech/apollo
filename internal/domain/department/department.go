package department

import "errors"

type Department struct {
	Id   int
	Name string // e.g. "ИУ-9", "ФН-12"
	Desc string // e.g. "Теоретическая информатика и компьютерные технологии"
}

var ErrNotExist = errors.New("department doesn't exist")
