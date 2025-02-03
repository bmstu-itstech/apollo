package material

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

type Material struct {
	Uuid       string
	Name       string
	Desc       string
	Url        string
	Author     string
	Views      int
	Department Department
	Discipline Discipline
	Created    time.Time
}

var ErrBadInput = errors.New("bad input")

var uuid_regex *regexp.Regexp = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

func NewMaterial(uuid, name, desc, url, author string, views int,
	dept Department, disc Discipline, t time.Time) (Material, error) {
	if !uuid_regex.Match([]byte(uuid)) {
		return Material{}, fmt.Errorf("%w (%s)", ErrBadInput, "uuid")
	}

	if name == "" || url == "" {
		// TODO: we could *theoretically* validate more here
		return Material{}, fmt.Errorf("%w (%s)", ErrBadInput, "name/url")
	}
	if (dept == Department{}) || (disc == Discipline{}) {
		// TODO: does this work?!
		return Material{}, fmt.Errorf("%w (%s)", ErrBadInput, "dept/disc")
	}

	m := Material{
		Uuid:       uuid,
		Name:       name,
		Desc:       desc, // no validation needed
		Url:        url,
		Author:     author, // no validation needed
		Views:      views,  // no validation needed
		Department: dept,
		Discipline: disc,
		Created:    t,
	}
	return m, nil
}
