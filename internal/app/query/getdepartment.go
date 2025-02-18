package query

import (
	"context"
	"log/slog"

	"github.com/bmstu-itstech/apollo/internal/common/decorator"
	"github.com/bmstu-itstech/apollo/internal/domain/material"
)

type GetDepartment struct {
	DepartmentId int
}

type GetDepartmentHandler decorator.QueryHandler[GetDepartment, Department]

type getDepartmentHandler struct {
	storage material.Storage
}

func NewGetDepartmentHandler(
	storage material.Storage,
	logger *slog.Logger,
) GetDepartmentHandler {
	if storage == nil {
		panic("storage is nil")
	}
	return decorator.ApplyQueryDecorators[GetDepartment, Department](
		getDepartmentHandler{storage: storage},
		logger,
	)
}

func (h getDepartmentHandler) Handle(ctx context.Context, query GetDepartment) (Department, error) {
	m, err := h.storage.Department(query.DepartmentId)
	if err != nil {
		return Department{}, err
	}
	return departmentFromDomain(m), nil
}
