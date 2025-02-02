package material

import (
	"time"

	"github.com/bmstu-itstech/apollo/internal/domain/department"
	"github.com/bmstu-itstech/apollo/internal/domain/discipline"
	"github.com/google/uuid"
)

type Material struct {
	Uuid       uuid.UUID `json:"uuid"`
	Name       string    `json:"name"`
	Desc       string    `json:"description"`
	Url        string    `json:"url"`
	Author     string    `json:"author"`
	Views      int       `json:"views"`
	Department department.Department
	Discipline discipline.Discipline
	Created    time.Time
}
