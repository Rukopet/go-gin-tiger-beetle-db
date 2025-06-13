package services

import (
	"context"

	"finance-api/internal/models"

	"github.com/google/uuid"
)

type AccountServiceInterface interface {
	CreateAccount(ctx context.Context, req *models.CreateAccountRequest) (*models.Account, error)
	GetAccount(ctx context.Context, id uuid.UUID) (*models.Account, error)
	GetAccounts(ctx context.Context, req *models.AccountsListRequest) (*models.PaginatedResponse, error)
	GetBalance(ctx context.Context, id uuid.UUID) (*models.BalanceResponse, error)
	Deposit(ctx context.Context, id uuid.UUID, req *models.DepositRequest) (*models.Transfer, error)
	Withdraw(ctx context.Context, id uuid.UUID, req *models.WithdrawRequest) (*models.Transfer, error)
	GetStatement(ctx context.Context, id uuid.UUID, dateFrom, dateTo, format string) (*models.StatementResponse, error)
}

type TransferServiceInterface interface {
	CreateTransfer(ctx context.Context, req *models.CreateTransferRequest) (*models.Transfer, error)
	GetTransfer(ctx context.Context, id uuid.UUID) (*models.Transfer, error)
	GetTransfers(ctx context.Context, req *models.TransfersListRequest) (*models.PaginatedResponse, error)
}

type StatsServiceInterface interface {
	GetSummary(ctx context.Context) (*models.StatsSummaryResponse, error)
}
