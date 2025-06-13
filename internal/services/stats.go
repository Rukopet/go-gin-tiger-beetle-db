package services

import (
	"context"

	"finance-api/internal/models"
)

type StatsService struct{}

func NewStatsService() *StatsService {
	return &StatsService{}
}

func (s *StatsService) GetSummary(ctx context.Context) (*models.StatsSummaryResponse, error) {
	return &models.StatsSummaryResponse{
		TotalAccounts:      0,
		TotalTransfers:     0,
		TotalTurnover:      0,
		AccountsByCurrency: map[string]int64{},
	}, nil
}
