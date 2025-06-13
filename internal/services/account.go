package services

import (
	"context"
	"time"

	"finance-api/internal/models"

	"github.com/google/uuid"
)

type AccountService struct{}

func NewAccountService() *AccountService {
	return &AccountService{}
}

func (s *AccountService) CreateAccount(ctx context.Context, req *models.CreateAccountRequest) (*models.Account, error) {
	return &models.Account{
		ID:        uuid.New(),
		Currency:  req.Currency,
		Balance:   req.InitialBalance,
		CreatedAt: time.Now(),
	}, nil
}

func (s *AccountService) GetAccount(ctx context.Context, id uuid.UUID) (*models.Account, error) {
	return &models.Account{
		ID:        id,
		Currency:  "USD",
		Balance:   0,
		CreatedAt: time.Now(),
	}, nil
}

func (s *AccountService) GetAccounts(ctx context.Context, req *models.AccountsListRequest) (*models.PaginatedResponse, error) {
	return &models.PaginatedResponse{
		Data:       []models.Account{},
		Page:       req.Page,
		Limit:      req.Limit,
		Total:      0,
		TotalPages: 0,
	}, nil
}

func (s *AccountService) GetBalance(ctx context.Context, id uuid.UUID) (*models.BalanceResponse, error) {
	return &models.BalanceResponse{
		AccountID: id.String(),
		Balance:   0,
		Currency:  "USD",
	}, nil
}

func (s *AccountService) Deposit(ctx context.Context, id uuid.UUID, req *models.DepositRequest) (*models.Transfer, error) {
	return &models.Transfer{
		ID:            uuid.New(),
		FromAccountID: uuid.New(),
		ToAccountID:   id,
		Amount:        req.Amount,
		Description:   req.Description,
		CreatedAt:     time.Now(),
		Status:        "completed",
	}, nil
}

func (s *AccountService) Withdraw(ctx context.Context, id uuid.UUID, req *models.WithdrawRequest) (*models.Transfer, error) {
	return &models.Transfer{
		ID:            uuid.New(),
		FromAccountID: id,
		ToAccountID:   uuid.New(),
		Amount:        req.Amount,
		Description:   req.Description,
		CreatedAt:     time.Now(),
		Status:        "completed",
	}, nil
}

func (s *AccountService) GetStatement(ctx context.Context, id uuid.UUID, dateFrom, dateTo, format string) (*models.StatementResponse, error) {
	return &models.StatementResponse{
		AccountID:    id.String(),
		DateFrom:     dateFrom,
		DateTo:       dateTo,
		OpenBalance:  0,
		CloseBalance: 0,
		Transactions: []models.TransactionRecord{},
	}, nil
}
