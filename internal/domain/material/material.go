package material

import (
	"encoding/json"
	"time"

	"github.com/bmstu-itstech/apollo/internal/domain/department"
	"github.com/bmstu-itstech/apollo/internal/domain/discipline"
)

type Material struct {
	Uuid       string                `json:"uuid"`
	Name       string                `json:"name"`
	Desc       string                `json:"description"`
	Url        string                `json:"url"`
	Author     string                `json:"author"`
	Views      int                   `json:"views"`
	Department department.Department
	Discipline discipline.Discipline
	Created    time.Time
}

func (m Material) JSONMarshal() ([]byte, error) {
	return json.Marshal(struct {
		Uuid       string `json:"uuid"`
		Name       string `json:"name"`
		Desc       string `json:"description"`
		Url        string `json:"url"`
		Author     string `json:"author"`
		Views      int    `json:"views"`
		Department int    `json:"department_id"`
		Discipline int    `json:"discipline_id"`
		Created    string `json:"created_at"`
	}{
		Uuid:       m.Uuid,
		Name:       m.Name,
		Desc:       m.Desc,
		Url:        m.Url,
		Author:     m.Author,
		Views:      m.Views,
		Department: m.Department.Id,
		Discipline: m.Discipline.Id,
		Created:    m.Created.Format(time.RFC3339),
	})
}
