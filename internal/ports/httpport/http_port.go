package httpport

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"


	"github.com/bmstu-itstech/apollo/internal/domain/material"
	"github.com/go-chi/render"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type Server struct {
	store material.Storage
}

// handleWrite "handles" error in output of ResponseWriter.Write.
func handleWrite(_ int, err error) {
	if err != nil {
		slog.Error("error writing response", "error", err)
	}
}

func (s Server) GetDepartaments(w http.ResponseWriter, r *http.Request) {
	departments, err := s.store.Departments()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	departments_bytes, err := json.Marshal(departments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	handleWrite(w.Write(departments_bytes))
}

func (s Server) GetDepartament(w http.ResponseWriter, r *http.Request, id int) {
	department, found := s.store.Department(id)
	if !found {
		http.Error(w, "department not found", http.StatusNotFound)
		return
	}
	department_bytes, err := json.Marshal(department)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	handleWrite(w.Write(department_bytes))
}

func (s Server) GetDisciplines(w http.ResponseWriter, r *http.Request) {
	disciplines, err := s.store.Disciplines()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	disciplines_bytes, err := json.Marshal(disciplines)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	handleWrite(w.Write(disciplines_bytes))
}

func (s Server) GetDiscipline(w http.ResponseWriter, r *http.Request, id int) {
	discipline, found := s.store.Discipline(id)
	if !found {
		http.Error(w, "discipline not found", http.StatusNotFound)
		return
	}
	discipline_bytes, err := json.Marshal(discipline)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	handleWrite(w.Write(discipline_bytes))
}

func (s Server) GetMaterials(w http.ResponseWriter, r *http.Request, params GetMaterialsParams) {
	discipline, found := s.store.Discipline(*params.DisciplineId)
	if !found {
		http.Error(w, "discipline doesn't exist", http.StatusBadRequest)
		return
	}
	materials, err := s.store.Materials(discipline)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	materials_bytes, err := json.Marshal(materials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	handleWrite(w.Write(materials_bytes))
}

func (s Server) GetMaterial(w http.ResponseWriter, r *http.Request, uuid openapi_types.UUID) {
	material, found := s.store.Material(uuid.String())
	if !found {
		http.Error(w, "material not found", http.StatusNotFound)
		return
	}
	material_bytes, err := json.Marshal(material)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	handleWrite(w.Write(material_bytes))
}

func (s Server) PutMaterial(w http.ResponseWriter, r *http.Request, uuid openapi_types.UUID) {
	var put_material PutMaterial
	if err := render.Decode(r, &put_material); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// TODO: Validate here?
	dept, found := s.store.Department(put_material.DepartamentId)
	if !found {
		http.Error(w, material.ErrDeptNotExist.Error(), http.StatusBadRequest)
		return
	}
	disc, found := s.store.Discipline(put_material.DisciplineId)
	if !found {
		http.Error(w, material.ErrDiscNotExist.Error(), http.StatusBadRequest)
		return
	}

	err := s.store.Upsert(uuid.String(), material.Material{
		Uuid:       uuid.String(),
		Name:       put_material.Name,
		Desc:       *put_material.Description,
		Url:        put_material.Url,
		Author:     *put_material.Author,
		Views:      0, // FIXME: Not part of PutMaterial?
		Department: dept,
		Discipline: disc,
		Created:    time.Now(), // FIXME: Same as views?
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	handleWrite(w.Write([]byte("OK")))
}


// func (m Material) JSONMarshal() ([]byte, error) {
// 	return json.Marshal(struct {
// 		Uuid       string `json:"uuid"`
// 		Name       string `json:"name"`
// 		Desc       string `json:"description"`
// 		Url        string `json:"url"`
// 		Author     string `json:"author"`
// 		Views      int    `json:"views"`
// 		Department int    `json:"department_id"`
// 		Discipline int    `json:"discipline_id"`
// 		Created    string `json:"created_at"`
// 	}{
// 		Uuid:       m.Uuid,
// 		Name:       m.Name,
// 		Desc:       m.Desc,
// 		Url:        m.Url,
// 		Author:     m.Author,
// 		Views:      m.Views,
// 		Department: m.Department.Id,
// 		Discipline: m.Discipline.Id,
// 		Created:    m.Created.Format(time.RFC3339),
// 	})
// }
