package command

import (
	"context"
	"log/slog"

	"github.com/bmstu-itstech/apollo/internal/common/decorator"
	"github.com/bmstu-itstech/apollo/internal/domain/material"
)

type Upsert struct {
	MaterialUUID string
	Material     material.Material
}

type UpsertHandler decorator.CommandHandler[Upsert]

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
	return decorator.ApplyCommandDecorators[Upsert](
		getMaterialHandler{storage: storage},
		logger,
	)
}

func (h getMaterialHandler) Handle(ctx context.Context, query Upsert) error {
	return h.storage.Upsert(query.MaterialUUID, query.Material)
}
