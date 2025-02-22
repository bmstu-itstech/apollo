package query

import (
	"context"
	"log/slog"

	"github.com/bmstu-itstech/apollo/internal/common/decorator"
	"github.com/bmstu-itstech/apollo/internal/domain/material"
)

type GetDiscipline struct {
	DisciplineId int
}

type GetDisciplineHandler decorator.QueryHandler[GetDiscipline, Discipline]

type getDisciplineHandler struct {
	storage material.Storage
}

func NewGetDisciplineHandler(
	storage material.Storage,
	logger *slog.Logger,
) GetDisciplineHandler {
	if storage == nil {
		panic("storage is nil")
	}
	return decorator.ApplyQueryDecorators[GetDiscipline, Discipline](
		getDisciplineHandler{storage: storage},
		logger,
	)
}

func (h getDisciplineHandler) Handle(ctx context.Context, query GetDiscipline) (Discipline, error) {
	d, err := h.storage.Discipline(query.DisciplineId)
	if err != nil {
		return Discipline{}, err
	}
	return disciplineFromDomain(d), nil
}
