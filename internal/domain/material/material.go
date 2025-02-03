package material

import (
	"errors"
	"regexp"
	"time"
)

type Material struct {
	Uuid         string
	Name         string
	Desc         string
	Url          string
	Author       *string
	Views        int
	DepartmentId int
	DisciplineId int
	Created      time.Time
}

var uuid_regex *regexp.Regexp = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

func UnmarshalMaterial(uuid, name, description, url string, author *string,
	views, department_id, discipline_id int, created_at time.Time) (Material, error) {
	if !uuid_regex.Match([]byte(uuid)) {
		return Material{}, errors.New("expected valid uuid")
	} else if name == "" {
		return Material{}, errors.New("expected non-empty name")
	} else if url == "" {
		return Material{}, errors.New("expected non-empty url")
	} else if department_id < 0 {
		return Material{}, errors.New("expected non-negative department id")
	} else if discipline_id < 0 {
		return Material{}, errors.New("expected non-negative discipline id")
	} else if views < 0 {
		return Material{}, errors.New("expected non-negative views")
	} else if author != nil && *author == "" {
		return Material{}, errors.New("expected non-empty author name or nil")
	}

	return Material{
		Uuid:         uuid,
		Name:         name,
		Desc:         description,
		Url:          url,
		Author:       author,
		Views:        views,
		DepartmentId: department_id,
		DisciplineId: discipline_id,
		Created:      created_at,
	}, nil
}

// NewMaterial creates a Material with default views and time
func NewMaterial(uuid, name, desc, url string, author *string, department, discipline int) (Material, error) {
	return UnmarshalMaterial(uuid, name, desc, url, author, 0, department, discipline, time.Now())
}

func MustNewMaterial(uuid, name, desc, url string, author *string, department, discipline int) Material {
	m, err := NewMaterial(uuid, name, desc, url, author, department, discipline)
	if err != nil {
		panic(err)
	}
	return m
}
