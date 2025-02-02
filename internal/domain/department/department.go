package department

import "errors"

type Department struct {
	Id   int    `json:"id"`
	Name string `json:"code"` // e.g. "ИУ-9", "ФН-12"
	Desc string `json:"name"` // e.g. "Теоретическая информатика и компьютерные технологии"
}

var ErrNotExist = errors.New("department doesn't exist")
