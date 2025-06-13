package models

import (
	"time"

	"github.com/google/uuid"
)

type Transfer struct {
	ID            uuid.UUID `json:"id"`
	FromAccountID uuid.UUID `json:"from_account_id"`
	ToAccountID   uuid.UUID `json:"to_account_id"`
	Amount        int64     `json:"amount"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
	Status        string    `json:"status"`
	TigerBeetleID uint128   `json:"tigerbeetle_id"`
}

type CreateTransferRequest struct {
	FromAccountID uuid.UUID `json:"from_account_id" binding:"required"`
	ToAccountID   uuid.UUID `json:"to_account_id" binding:"required"`
	Amount        int64     `json:"amount" binding:"required,min=1"`
	Description   string    `json:"description"`
}

type TransfersListRequest struct {
	Page      int       `form:"page,default=1" binding:"min=1"`
	Limit     int       `form:"limit,default=10" binding:"min=1,max=100"`
	AccountID uuid.UUID `form:"account_id,omitempty"`
	DateFrom  string    `form:"date_from,omitempty"`
	DateTo    string    `form:"date_to,omitempty"`
}
