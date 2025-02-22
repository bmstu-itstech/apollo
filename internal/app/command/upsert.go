package command

import (
	"context"
	"log/slog"

	"github.com/bmstu-itstech/apollo/internal/common/decorator"
	"github.com/bmstu-itstech/apollo/internal/domain/material"
)

type UpsertMaterial struct {
	UUID         string
	Name         string
	Desc         string
	Url          string
	Author       *string
	DepartmentId int
	DisciplineId int
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
	mat, err = material.NewMaterial(
		query.UUID,
		query.Name,
		query.Desc,
		query.Url,
		query.Author,
		query.DepartmentId,
		query.DisciplineId,
	)
	if err != nil {
		return err
	}
	return h.storage.Upsert(mat.UUID, mat)
}
