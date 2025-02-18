package query

import (
	"context"
	"log/slog"

	"github.com/bmstu-itstech/apollo/internal/common/decorator"
	"github.com/bmstu-itstech/apollo/internal/domain/material"
)

type GetDisciplines struct {
}

type GetDisciplinesHandler decorator.QueryHandler[GetDisciplines, []Discipline]

type getDisciplinesHandler struct {
	storage material.Storage
}

func NewGetDisciplinesHandler(
	storage material.Storage,
	logger *slog.Logger,
) GetDisciplinesHandler {
	if storage == nil {
		panic("storage is nil")
	}
	return decorator.ApplyQueryDecorators[GetDisciplines, []Discipline](
		getDisciplinesHandler{storage: storage},
		logger,
	)
}

func (h getDisciplinesHandler) Handle(ctx context.Context, query GetDisciplines) ([]Discipline, error) {
	mat, err := h.storage.Disciplines()
	if err != nil {
		return nil, err
	}
	mats_fd := make([]Discipline, 0)
	for _, m := range mat {
		mats_fd = append(mats_fd, disciplineFromDomain(m))
	}
	return mats_fd, nil
}
