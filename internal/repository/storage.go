package repository

import (
	"context"

	"github.com/arvaliullin/vhagar/internal/core/domain"
)

//go:generate mockgen -source=storage.go -destination=mock/storage_mock.go -package=repomock

type MetricStorage interface {
	Ping(ctx context.Context) error
	Save(ctx context.Context, metric *domain.Metric) error
	GetByID(ctx context.Context, id string) (*domain.Metric, error)
	GetAll(ctx context.Context) ([]*domain.Metric, error)
	Delete(ctx context.Context, id string) error
}

