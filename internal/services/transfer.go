package services

import (
	"context"
	"time"

	"finance-api/internal/models"

	"github.com/google/uuid"
)

type TransferService struct{}

func NewTransferService() *TransferService {
	return &TransferService{}
}

func (s *TransferService) CreateTransfer(ctx context.Context, req *models.CreateTransferRequest) (*models.Transfer, error) {
	return &models.Transfer{
		ID:            uuid.New(),
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
		Description:   req.Description,
		CreatedAt:     time.Now(),
		Status:        "completed",
	}, nil
}

func (s *TransferService) GetTransfer(ctx context.Context, id uuid.UUID) (*models.Transfer, error) {
	return &models.Transfer{
		ID:            id,
		FromAccountID: uuid.New(),
		ToAccountID:   uuid.New(),
		Amount:        100,
		Description:   "Sample transfer",
		CreatedAt:     time.Now(),
		Status:        "completed",
	}, nil
}

func (s *TransferService) GetTransfers(ctx context.Context, req *models.TransfersListRequest) (*models.PaginatedResponse, error) {
	return &models.PaginatedResponse{
		Data:       []models.Transfer{},
		Page:       req.Page,
		Limit:      req.Limit,
		Total:      0,
		TotalPages: 0,
	}, nil
}
