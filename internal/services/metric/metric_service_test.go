package metric

import (
	"context"
	"testing"

	"github.com/arvaliullin/vhagar/internal/core/domain"
	repomock "github.com/arvaliullin/vhagar/internal/repository/mock"
	"github.com/golang/mock/gomock"
)

func TestService_CreateMetric(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := repomock.NewMockMetricStorage(ctrl)
	service := NewService(mockStorage)

	tests := []struct {
		name    string
		metric  *domain.Metric
		setup   func()
		wantErr bool
	}{
		{
			name: "successful creation",
			metric: &domain.Metric{
				ID:    "test-id",
				Name:  "test-metric",
				Value: 42.0,
				Type:  "gauge",
			},
			setup: func() {
				mockStorage.EXPECT().
					Save(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			wantErr: false,
		},
		{
			name:    "nil metric",
			metric:  nil,
			setup:   func() {},
			wantErr: true,
		},
		{
			name: "empty ID",
			metric: &domain.Metric{
				ID:    "",
				Name:  "test-metric",
				Value: 42.0,
				Type:  "gauge",
			},
			setup:   func() {},
			wantErr: true,
		},
		{
			name: "empty Name",
			metric: &domain.Metric{
				ID:    "test-id",
				Name:  "",
				Value: 42.0,
				Type:  "gauge",
			},
			setup:   func() {},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			err := service.CreateMetric(context.Background(), tt.metric)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateMetric() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_GetMetric(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := repomock.NewMockMetricStorage(ctrl)
	service := NewService(mockStorage)

	tests := []struct {
		name    string
		id      string
		setup   func()
		wantErr bool
	}{
		{
			name: "successful get",
			id:   "test-id",
			setup: func() {
				mockStorage.EXPECT().
					GetByID(gomock.Any(), "test-id").
					Return(&domain.Metric{
						ID:    "test-id",
						Name:  "test-metric",
						Value: 42.0,
						Type:  "gauge",
					}, nil)
			},
			wantErr: false,
		},
		{
			name:    "empty ID",
			id:      "",
			setup:   func() {},
			wantErr: true,
		},
		{
			name: "not found",
			id:   "non-existent",
			setup: func() {
				mockStorage.EXPECT().
					GetByID(gomock.Any(), "non-existent").
					Return(nil, nil)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			_, err := service.GetMetric(context.Background(), tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMetric() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_ListMetrics(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := repomock.NewMockMetricStorage(ctrl)
	service := NewService(mockStorage)

	mockStorage.EXPECT().
		GetAll(gomock.Any()).
		Return([]*domain.Metric{
			{
				ID:    "id1",
				Name:  "metric1",
				Value: 10.0,
				Type:  "gauge",
			},
			{
				ID:    "id2",
				Name:  "metric2",
				Value: 20.0,
				Type:  "counter",
			},
		}, nil)

	metrics, err := service.ListMetrics(context.Background())
	if err != nil {
		t.Errorf("ListMetrics() error = %v", err)
	}

	if len(metrics) != 2 {
		t.Errorf("ListMetrics() returned %d metrics, want 2", len(metrics))
	}
}

func TestService_DeleteMetric(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := repomock.NewMockMetricStorage(ctrl)
	service := NewService(mockStorage)

	tests := []struct {
		name    string
		id      string
		setup   func()
		wantErr bool
	}{
		{
			name: "successful delete",
			id:   "test-id",
			setup: func() {
				mockStorage.EXPECT().
					Delete(gomock.Any(), "test-id").
					Return(nil)
			},
			wantErr: false,
		},
		{
			name:    "empty ID",
			id:      "",
			setup:   func() {},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			err := service.DeleteMetric(context.Background(), tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteMetric() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_HealthCheck(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := repomock.NewMockMetricStorage(ctrl)
	service := NewService(mockStorage)

	mockStorage.EXPECT().
		Ping(gomock.Any()).
		Return(nil)

	err := service.HealthCheck(context.Background())
	if err != nil {
		t.Errorf("HealthCheck() error = %v", err)
	}
}

