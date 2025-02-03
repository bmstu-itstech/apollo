package sqlite

import (
	"database/sql"
	"time"

	"github.com/bmstu-itstech/apollo/internal/domain/material"
	_ "github.com/mattn/go-sqlite3" // todo: _ -> sqlite?
)

// SQLiteStorage is a Storage implementation using SQLite3.
// It is not intended for production, as separate DBs are better fit for that.
// However, it may be useful for local development
type SQLiteStorage struct {
	db *sql.DB
}

func (s SQLiteStorage) Department(id int) (material.Department, bool) {
	row := s.db.QueryRow("select * from departments where id = $1", id)
	department := material.Department{}
	err := row.Scan(&department.Id, &department.Name, &department.Desc)
	return department, err == nil
}

func (s SQLiteStorage) Departments() ([]material.Department, error) {
	rows, err := s.db.Query("select * from departments")
	if err != nil {
		return nil, err
	}
	departments := make([]material.Department, 0)
	for rows.Next() {
		department := material.Department{}
		err = rows.Scan(&department.Id, &department.Name, &department.Desc)
		if err != nil {
			return nil, err
		}
		departments = append(departments, department)
	}
	return departments, nil
}

func (s SQLiteStorage) Discipline(id int) (material.Discipline, bool) {
	row := s.db.QueryRow("select * from disciplines where id = $1", id)
	discipline := material.Discipline{}
	err := row.Scan(&discipline.Id, &discipline.Name)
	return discipline, err == nil
}

func scanMaterial(row *sql.Rows, s SQLiteStorage) (material.Material, error) {
	m := material.Material{}
	var department_id int
	var discipline_id int
	var created_at int64

	err := row.Scan(&m.Uuid, &m.Name, &m.Desc,
		&m.Url, &m.Author, &m.Views,
		&department_id, &discipline_id, &created_at)
	if err != nil {
		return m, err
	}
	m_department, found := s.Department(department_id)
	if !found {
		return m, material.ErrDeptNotExist
	}
	m.Department = m_department

	m_discipline, found := s.Discipline(discipline_id)
	if !found {
		return m, material.ErrDiscNotExist
	}
	m.Discipline = m_discipline

	m.Created = time.Unix(created_at, 0)
	return m, nil
}

func (s SQLiteStorage) Materials(discipline material.Discipline) ([]material.Material, error) {
	rows, err := s.db.Query("select * from materials where discipline_id = $1", discipline.Id)
	if err != nil {
		return nil, err
	}
	materials := make([]material.Material, 0)
	for rows.Next() {
		material, err := scanMaterial(rows, s)
		if err != nil {
			return nil, err
		}
		materials = append(materials, material)
	}
	return materials, nil
}

func (s SQLiteStorage) Material(uuid string) (material.Material, bool) {
	// TODO: the bool feels like potentially hiding errors under the rug
	rows, err := s.db.Query("select * from materials where uuid = $1", uuid)
	if err != nil || !rows.Next() {
		return material.Material{}, false
	}
	material, err := scanMaterial(rows, s)
	return material, err == nil
}

func (s SQLiteStorage) Upsert(uuid string, material material.Material) error {
	_, err := s.db.Exec(`
	insert into materials
		(uuid, name, description, url,
		author, views, department_id,
		discipline_id, created_at)
	values ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		material.Uuid, material.Name, material.Desc, material.Url,
		material.Author, material.Views, material.Department.Id,
		material.Discipline.Id, material.Created.Unix())
	return err
}

// TODO: Disciplines, DeptDisciplines
