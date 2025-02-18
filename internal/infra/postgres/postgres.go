package postgres

import (
	"context"
	"errors"

	"github.com/bmstu-itstech/apollo/internal/domain/material"
	"github.com/jackc/pgx/v5"
)

// PgStorage is a storage implementation using Postgres
type PgStorage struct {
	conn *pgx.Conn
}

func NewPgStorage(conn *pgx.Conn) PgStorage {
	return PgStorage{conn: conn}
}

func (s PgStorage) Materials(discipline_id int) ([]material.Material, error) {
	rows, err := s.conn.Query(context.Background(), "select * from materials where discipline_id = $1", discipline_id)
	if errors.Is(err, pgx.ErrNoRows) {
		return []material.Material{}, nil // no relevant materials exist
	} else if err != nil {
		return nil, err
	}
	matRows, err := pgx.CollectRows(rows, pgx.RowToStructByName[materialRow])
	if err != nil {
		return nil, err
	}
	mats := make([]material.Material, 0, len(matRows))
	for _, m := range matRows {
		mm, err := m.mapFrom()
		if err != nil {
			return nil, err
		}
		mats = append(mats, mm)
	}
	return mats, nil
}

func (s PgStorage) Material(uuid string) (material.Material, error) {
	rows, err := s.conn.Query(context.Background(), "select * from materials where uuid = $1", uuid)
	if errors.Is(err, pgx.ErrNoRows) {
		return material.Material{}, material.ErrMatNotExist
	} else if err != nil {
		return material.Material{}, err
	}
	mat, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[materialRow])
	if err != nil {
		return material.Material{}, err
	}
	return mat.mapFrom()
}

func (s PgStorage) Upsert(uuid string, material material.Material) error {
	// note how on conflict, we don't touch views and created_at
	_, err := s.conn.Exec(context.Background(), `
	insert into materials
		(uuid, name, description, url,
		author, views, department_id,
		discipline_id, created_at)
	values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
    on conflict (uuid) do update set
        uuid = EXCLUDED.uuid,
        name = EXCLUDED.name,
        description = EXCLUDED.description,
        url = EXCLUDED.url,
        author = EXCLUDED.author,
        department_id = EXCLUDED.department_id,
        discipline_id = EXCLUDED.discipline_id,
    `,
		material.UUID, material.Name, material.Desc, material.Url,
		material.Author, material.Views, material.DepartmentId,
		material.DisciplineId, material.Created)
	return err
}

func (s PgStorage) Disciplines() ([]material.Discipline, error) {
	rows, err := s.conn.Query(context.Background(), "select * from disciplines")
	if errors.Is(err, pgx.ErrNoRows) {
		return []material.Discipline{}, nil // no disciplines exist
	} else if err != nil {
		return nil, err
	}
	discRows, err := pgx.CollectRows(rows, pgx.RowToStructByName[disciplineRow])
	if err != nil {
		return nil, err
	}
	discs := make([]material.Discipline, 0, len(discRows))
	for _, d := range discRows {
		dd, err := d.mapFrom()
		if err != nil {
			return nil, err
		}
		discs = append(discs, dd)
	}
	return discs, nil
}

func (s PgStorage) Discipline(id int) (material.Discipline, error) {
	rows, err := s.conn.Query(context.Background(), "select * from disciplines where id = $1", id)
	if errors.Is(err, pgx.ErrNoRows) {
		return material.Discipline{}, material.ErrDiscNotExist
	} else if err != nil {
		return material.Discipline{}, err
	}
	disc, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[disciplineRow])
	if err != nil {
		return material.Discipline{}, err
	}
	return disc.mapFrom()
}

func (s PgStorage) Departments() ([]material.Department, error) {
	rows, err := s.conn.Query(context.Background(), "select * from departments")
	if errors.Is(err, pgx.ErrNoRows) {
		return []material.Department{}, nil // no departments exist
	} else if err != nil {
		return nil, err
	}
	depRows, err := pgx.CollectRows(rows, pgx.RowToStructByName[departmentRow])
	if err != nil {
		return nil, err
	}
	deps := make([]material.Department, 0, len(depRows))
	for _, d := range depRows {
		dd, err := d.mapFrom()
		if err != nil {
			return nil, err
		}
		deps = append(deps, dd)
	}
	return deps, nil
}

func (s PgStorage) Department(id int) (material.Department, error) {
	rows, err := s.conn.Query(context.Background(), "select * from departments where id = $1", id)
	if errors.Is(err, pgx.ErrNoRows) {
		return material.Department{}, material.ErrDeptNotExist
	} else if err != nil {
		return material.Department{}, err
	}
	dep, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[departmentRow])
	if err != nil {
		return material.Department{}, err
	}
	return dep.mapFrom()
}
