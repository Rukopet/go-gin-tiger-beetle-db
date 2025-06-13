package models

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID            uuid.UUID `json:"id"`
	Currency      string    `json:"currency"`
	Balance       int64     `json:"balance"`
	CreatedAt     time.Time `json:"created_at"`
	TigerBeetleID uint128   `json:"tigerbeetle_id"`
}

type CreateAccountRequest struct {
	Currency       string `json:"currency" binding:"required,oneof=USD EUR RUB"`
	InitialBalance int64  `json:"initial_balance,omitempty"`
}

type DepositRequest struct {
	Amount      int64  `json:"amount" binding:"required,min=1"`
	Description string `json:"description"`
}

type WithdrawRequest struct {
	Amount      int64  `json:"amount" binding:"required,min=1"`
	Description string `json:"description"`
}

type AccountsListRequest struct {
	Page     int    `form:"page,default=1" binding:"min=1"`
	Limit    int    `form:"limit,default=10" binding:"min=1,max=100"`
	Currency string `form:"currency,omitempty" binding:"omitempty,oneof=USD EUR RUB"`
}

type uint128 [16]byte
