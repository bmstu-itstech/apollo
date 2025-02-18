package query

import (
	"context"
	"log/slog"

	"github.com/bmstu-itstech/apollo/internal/common/decorator"
	"github.com/bmstu-itstech/apollo/internal/domain/material"
)

type GetMaterials struct {
	DisciplineId int
}

type GetMaterialsHandler decorator.QueryHandler[GetMaterials, []Material]

type getMaterialsHandler struct {
	storage material.Storage
}

func NewGetMaterialsHandler(
	storage material.Storage,
	logger *slog.Logger,
) GetMaterialsHandler {
	if storage == nil {
		panic("storage is nil")
	}
	return decorator.ApplyQueryDecorators[GetMaterials, []Material](
		getMaterialsHandler{storage: storage},
		logger,
	)
}

func (h getMaterialsHandler) Handle(ctx context.Context, query GetMaterials) ([]Material, error) {
	mats, err := h.storage.Materials(query.DisciplineId)
	if err != nil {
		return nil, err
	}
	return materialsFromDomain(mats), nil
}
