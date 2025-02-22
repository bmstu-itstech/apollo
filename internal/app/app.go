package app

import (
	"github.com/bmstu-itstech/apollo/internal/app/command"
	"github.com/bmstu-itstech/apollo/internal/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	UpsertMaterial command.UpsertHandler
}

type Queries struct {
	GetMaterials   query.GetMaterialsHandler
	GetMaterial    query.GetMaterialHandler
	GetDisciplines query.GetDisciplinesHandler
	GetDiscipline  query.GetDisciplineHandler
	GetDepartments query.GetDepartmentsHandler
	GetDepartment  query.GetDepartmentHandler
}
