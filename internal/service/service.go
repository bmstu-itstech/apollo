package service

import (
	"log/slog"

	"github.com/bmstu-itstech/apollo/internal/app"
	"github.com/bmstu-itstech/apollo/internal/app/command"
	"github.com/bmstu-itstech/apollo/internal/app/query"
	"github.com/bmstu-itstech/apollo/internal/common/logs"
	"github.com/bmstu-itstech/apollo/internal/domain/material"
	mock_infra "github.com/bmstu-itstech/apollo/internal/infra/mock"
)

type Cleanup func()

func NewApplication() (*app.Application, Cleanup) {
	logger := logs.DefaultLogger()

	// url := os.Getenv("DATABASE_URI")
	// conn, err := pgx.Connect(context.Background(), url)
	// if err != nil {
	// 	panic(err)
	// }
	// store := postgres.NewPgStorage(conn)
	store := mock_infra.NewMockStorage()

	return newApplication(logger, store), func() {
		// _ = conn.Close(context.Background())
	}
}

func NewTestApplication() *app.Application {
	logger := logs.DefaultLogger()
	store := mock_infra.NewMockStorage()
	return newApplication(logger, store)
}

func newApplication(
	logger *slog.Logger,
	store material.Storage,
) *app.Application {
	return &app.Application{
		Commands: app.Commands{
			UpsertMaterial: command.NewUpsertHandler(store, logger),
		},
		Queries: app.Queries{
			GetMaterials:   query.NewGetMaterialsHandler(store, logger),
			GetMaterial:    query.NewGetMaterialHandler(store, logger),
			GetDisciplines: query.NewGetDisciplinesHandler(store, logger),
			GetDiscipline:  query.NewGetDisciplineHandler(store, logger),
			GetDepartments: query.NewGetDepartmentsHandler(store, logger),
			GetDepartment:  query.NewGetDepartmentHandler(store, logger),
		},
	}
}
