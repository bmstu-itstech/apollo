package query

import (
	"context"
	"log/slog"

	"github.com/bmstu-itstech/apollo/internal/common/decorator"
	"github.com/bmstu-itstech/apollo/internal/domain/material"
)

type GetMaterial struct {
	MaterialUUID string
}

type GetMaterialHandler decorator.QueryHandler[GetMaterial, Material]

type getMaterialHandler struct {
	storage material.Storage
}

func NewGetMaterialHandler(
	storage material.Storage,
	logger *slog.Logger,
) GetMaterialHandler {
	if storage == nil {
		panic("storage is nil")
	}
	return decorator.ApplyQueryDecorators[GetMaterial, Material](
		getMaterialHandler{storage: storage},
		logger,
	)
}

func (h getMaterialHandler) Handle(ctx context.Context, query GetMaterial) (Material, error) {
	m, err := h.storage.Material(query.MaterialUUID)
	if err != nil {
		return Material{}, err
	}
	return materialFromDomain(m), nil
}
