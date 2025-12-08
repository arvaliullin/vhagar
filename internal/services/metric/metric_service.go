package metric

import (
	"context"
	"errors"

	"github.com/arvaliullin/vhagar/internal/core/domain"
	"github.com/arvaliullin/vhagar/internal/repository"
)

var (
	ErrMetricNotFound = errors.New("metric not found")
	ErrInvalidMetric  = errors.New("invalid metric")
)

type Service struct {
	storage repository.MetricStorage
}

func NewService(storage repository.MetricStorage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) CreateMetric(ctx context.Context, metric *domain.Metric) error {
	if metric == nil || metric.ID == "" || metric.Name == "" {
		return ErrInvalidMetric
	}

	return s.storage.Save(ctx, metric)
}

func (s *Service) GetMetric(ctx context.Context, id string) (*domain.Metric, error) {
	if id == "" {
		return nil, ErrInvalidMetric
	}

	metric, err := s.storage.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if metric == nil {
		return nil, ErrMetricNotFound
	}

	return metric, nil
}

func (s *Service) ListMetrics(ctx context.Context) ([]*domain.Metric, error) {
	return s.storage.GetAll(ctx)
}

func (s *Service) DeleteMetric(ctx context.Context, id string) error {
	if id == "" {
		return ErrInvalidMetric
	}

	return s.storage.Delete(ctx, id)
}

func (s *Service) HealthCheck(ctx context.Context) error {
	return s.storage.Ping(ctx)
}
