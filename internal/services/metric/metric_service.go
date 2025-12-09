// Package metric предоставляет бизнес-логику для управления метриками.
package metric

import (
	"context"
	"errors"

	"github.com/arvaliullin/vhagar/internal/core/domain"
	"github.com/arvaliullin/vhagar/internal/repository"
)

var (
	// ErrMetricNotFound возвращается, когда запрашиваемая метрика не найдена.
	ErrMetricNotFound = errors.New("metric not found")
	// ErrInvalidMetric возвращается, когда метрика не проходит валидацию.
	ErrInvalidMetric = errors.New("invalid metric")
)

// Service предоставляет бизнес-логику для операций с метриками.
type Service struct {
	storage repository.MetricStorage
}

// NewService создает новый сервис метрик с указанным хранилищем.
func NewService(storage repository.MetricStorage) *Service {
	return &Service{
		storage: storage,
	}
}

// CreateMetric создает новую метрику после валидации.
func (s *Service) CreateMetric(ctx context.Context, metric *domain.Metric) error {
	if metric == nil || metric.ID == "" || metric.Name == "" {
		return ErrInvalidMetric
	}

	return s.storage.Save(ctx, metric)
}

// GetMetric получает метрику по её идентификатору.
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

// ListMetrics получает все метрики.
func (s *Service) ListMetrics(ctx context.Context) ([]*domain.Metric, error) {
	return s.storage.GetAll(ctx)
}

// DeleteMetric удаляет метрику по её идентификатору.
func (s *Service) DeleteMetric(ctx context.Context, id string) error {
	if id == "" {
		return ErrInvalidMetric
	}

	return s.storage.Delete(ctx, id)
}

// HealthCheck выполняет проверку работоспособности хранилища.
func (s *Service) HealthCheck(ctx context.Context) error {
	return s.storage.Ping(ctx)
}
