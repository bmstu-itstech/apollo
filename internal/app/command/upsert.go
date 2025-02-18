package command

import (
	"context"
	"log/slog"
	"time"

	"github.com/bmstu-itstech/apollo/internal/common/decorator"
	"github.com/bmstu-itstech/apollo/internal/domain/material"
)

type UpsertMaterial struct {
	UUID         string
	Name         string
	Desc         string
	Url          string
	Author       *string
	Views        *int
	DepartmentId int
	DisciplineId int
	Created      *time.Time
}

type UpsertHandler decorator.CommandHandler[UpsertMaterial]

type getMaterialHandler struct {
	storage material.Storage
}

func NewUpsertHandler(
	storage material.Storage,
	logger *slog.Logger,
) UpsertHandler {
	if storage == nil {
		panic("storage is nil")
	}
	return decorator.ApplyCommandDecorators[UpsertMaterial](
		getMaterialHandler{storage: storage},
		logger,
	)
}

func (h getMaterialHandler) Handle(ctx context.Context, query UpsertMaterial) error {
	var mat material.Material
	var err error
	if query.Views == nil || query.Created == nil { // Adding new material
		mat, err = material.NewMaterial(
			query.UUID,
			query.Name,
			query.Desc,
			query.Url,
			query.Author,
			query.DepartmentId,
			query.DisciplineId,
		)
	} else {
		mat, err = material.UnmarshalMaterial(
			query.UUID,
			query.Name,
			query.Desc,
			query.Url,
			query.Author,
			*query.Views,
			query.DepartmentId,
			query.DisciplineId,
			*query.Created,
		)
	}
	if err != nil {
		return err
	}
	return h.storage.Upsert(mat.UUID, mat)
}
