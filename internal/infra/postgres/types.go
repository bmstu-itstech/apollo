package postgres

import (
	"time"

	"github.com/bmstu-itstech/apollo/internal/domain/material"
)

type materialRow struct {
	UUID         string    `db:"uuid"`
	Name         string    `db:"name"`
	Desc         string    `db:"description"`
	Url          string    `db:"url"`
	Author       *string   `db:"author"`
	Views        int       `db:"views"`
	DepartmentId int       `db:"department_id"`
	DisciplineId int       `db:"discipline_id"`
	Created      time.Time `db:"created_at"`
}

func (m materialRow) mapFrom() (material.Material, error) {
	return material.UnmarshalMaterial(m.UUID, m.Name, m.Desc, m.Url, m.Author,
		m.Views, m.DepartmentId, m.DisciplineId, m.Created)
}

func mapMaterialToRow(m material.Material) materialRow {
	return materialRow{
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

type departmentRow struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Desc string `db:"description"`
}

func (d departmentRow) mapFrom() (material.Department, error) {
	return material.NewDepartment(d.Id, d.Name, d.Desc)
}

func mapDepartmentToRow(d material.Department) departmentRow {
	return departmentRow{
		Id:   d.Id,
		Name: d.Name,
		Desc: d.Desc,
	}
}

type disciplineRow struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Desc string `db:"description"`
}

func (d disciplineRow) mapFrom() (material.Discipline, error) {
	return material.NewDiscipline(d.Id, d.Name)
}

func mapDisciplineToRow(d material.Discipline) disciplineRow {
	return disciplineRow{
		Id:   d.Id,
		Name: d.Name,
	}
}
