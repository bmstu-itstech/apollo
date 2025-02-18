package query

import (
	"time"

	"github.com/bmstu-itstech/apollo/internal/domain/material"
)

type Material struct {
	UUID         string
	Name         string
	Desc         string
	Url          string
	Author       *string
	Views        int
	DepartmentId int
	DisciplineId int
	Created      time.Time
}

func materialFromDomain(m material.Material) Material {
	return Material{
		UUID:         m.UUID,
		Name:         m.Name,
		Desc:         m.Desc,
		Url:          m.Url,
		Author:       m.Author,
		Views:        m.Views,
		DepartmentId: m.DepartmentId,
		DisciplineId: m.DisciplineId,
		Created:      m.Created,
	}
}

type Department struct {
	Id   int
	Name string
	Desc string
}

func departmentFromDomain(d material.Department) Department {
	return Department{
		Id:   d.Id,
		Name: d.Name,
		Desc: d.Desc,
	}
}

type Discipline struct {
	Id   int
	Name string
}

func disciplineFromDomain(d material.Discipline) Discipline {
	return Discipline{
		Id:   d.Id,
		Name: d.Name,
	}
}
