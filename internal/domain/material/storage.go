package material

type Storage interface {
	Materials(discipline Discipline) ([]Material, error)
	Material(uuid string) (Material, bool)
	Upsert(uuid string, material Material) error
	Disciplines() ([]Discipline, error)
	Discipline(id int) (Discipline, bool)
	Departments() ([]Department, error)
	Department(id int) (Department, bool)
}
