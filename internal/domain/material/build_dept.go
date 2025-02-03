package material

import (
	"encoding/json"
	"io"
	"net/http"
)

// This file is mostly provided for fun and future reference.
// It attempts to create a list of all departments through an external api.
// An issue we encounter there is that you can't map UUIDs to integers.

type StructureData struct {
	Data StructureNode `json:"data"`
	Date string        `json:"date"` // likely useless
}
type StructureNode struct {
	UUID         string          `json:"uuid"`
	Abbreviation string          `json:"abbr"`
	Name         string          `json:"name"`
	Children     []StructureNode `json:"children"`
	Type         string          `json:"nodeType"`
}

// BuildDepartments iterates over lks.bmstu.ru's "structure" api and finds all relevant nodes
func BuildDepartments() ([]Department, error) {
	response, err := http.Get("https://lks.bmstu.ru/lks-back/api/v1/structure")
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		// what error? custom?
		return nil, nil
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	structure := StructureData{}
	err = json.Unmarshal(responseBody, &structure)
	if err != nil {
		return nil, err
	}

	departments := make([]Department, 0)
	i := 0
	var parse func(node StructureNode)
	parse = func(node StructureNode) {
		if node.Type == "department" {
			// node.UUID -> ???
			departments = append(departments,
				Department{i, node.Abbreviation, node.Name})
			i++
		}
		for _, child := range node.Children {
			parse(child)
		}
	}
	parse(structure.Data)

	return departments, nil
}
