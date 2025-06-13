package models

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
}

type SuccessResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message,omitempty"`
}

type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	Total      int64       `json:"total"`
	TotalPages int         `json:"total_pages"`
}

type BalanceResponse struct {
	AccountID string `json:"account_id"`
	Balance   int64  `json:"balance"`
	Currency  string `json:"currency"`
}

type StatsSummaryResponse struct {
	TotalAccounts      int64            `json:"total_accounts"`
	TotalTransfers     int64            `json:"total_transfers"`
	TotalTurnover      int64            `json:"total_turnover"`
	AccountsByCurrency map[string]int64 `json:"accounts_by_currency"`
}

type StatementResponse struct {
	AccountID    string              `json:"account_id"`
	DateFrom     string              `json:"date_from"`
	DateTo       string              `json:"date_to"`
	OpenBalance  int64               `json:"open_balance"`
	CloseBalance int64               `json:"close_balance"`
	Transactions []TransactionRecord `json:"transactions"`
}

type TransactionRecord struct {
	ID          string `json:"id"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Amount      int64  `json:"amount"`
	Type        string `json:"type"`
	Balance     int64  `json:"balance"`
}
