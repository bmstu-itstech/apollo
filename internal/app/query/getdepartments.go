package query

import (
	"context"
	"log/slog"

	"github.com/bmstu-itstech/apollo/internal/common/decorator"
	"github.com/bmstu-itstech/apollo/internal/domain/material"
)

type GetDepartments struct {
}

type GetDepartmentsHandler decorator.QueryHandler[GetDepartments, []Department]

type getDepartmentsHandler struct {
	storage material.Storage
}

func NewGetDepartmentsHandler(
	storage material.Storage,
	logger *slog.Logger,
) GetDepartmentsHandler {
	if storage == nil {
		panic("storage is nil")
	}
	return decorator.ApplyQueryDecorators[GetDepartments, []Department](
		getDepartmentsHandler{storage: storage},
		logger,
	)
}

func (h getDepartmentsHandler) Handle(ctx context.Context, query GetDepartments) ([]Department, error) {
	deps, err := h.storage.Departments()
	if err != nil {
		return nil, err
	}
	return departmentsFromDomain(deps), nil
}
