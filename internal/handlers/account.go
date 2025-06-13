package handlers

import (
	"net/http"

	"finance-api/internal/models"
	"finance-api/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AccountHandler struct {
	accountService *services.AccountService
}

func NewAccountHandler(accountService *services.AccountService) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
	}
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {
	var req models.CreateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Invalid request",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	account, err := h.accountService.CreateAccount(c.Request.Context(), &req)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusCreated, models.SuccessResponse{
		Data:    account,
		Message: "Account created successfully",
	})
}

func (h *AccountHandler) GetAccount(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Invalid ID",
			Message: "Invalid UUID format",
			Code:    http.StatusBadRequest,
		})
		return
	}

	account, err := h.accountService.GetAccount(c.Request.Context(), id)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Data: account,
	})
}

func (h *AccountHandler) GetAccounts(c *gin.Context) {
	var req models.AccountsListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Invalid query parameters",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	if req.Page == 0 {
		req.Page = 1
	}
	if req.Limit == 0 {
		req.Limit = 10
	}

	response, err := h.accountService.GetAccounts(c.Request.Context(), &req)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *AccountHandler) GetBalance(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Invalid ID",
			Message: "Invalid UUID format",
			Code:    http.StatusBadRequest,
		})
		return
	}

	balance, err := h.accountService.GetBalance(c.Request.Context(), id)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Data: balance,
	})
}

func (h *AccountHandler) Deposit(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Invalid ID",
			Message: "Invalid UUID format",
			Code:    http.StatusBadRequest,
		})
		return
	}

	var req models.DepositRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Invalid request",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	transfer, err := h.accountService.Deposit(c.Request.Context(), id, &req)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Data:    transfer,
		Message: "Deposit processed successfully",
	})
}

func (h *AccountHandler) Withdraw(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Invalid ID",
			Message: "Invalid UUID format",
			Code:    http.StatusBadRequest,
		})
		return
	}

	var req models.WithdrawRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Invalid request",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	transfer, err := h.accountService.Withdraw(c.Request.Context(), id, &req)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Data:    transfer,
		Message: "Withdrawal processed successfully",
	})
}

func (h *AccountHandler) GetStatement(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Invalid ID",
			Message: "Invalid UUID format",
			Code:    http.StatusBadRequest,
		})
		return
	}

	dateFrom := c.Query("date_from")
	dateTo := c.Query("date_to")
	format := c.DefaultQuery("format", "JSON")

	statement, err := h.accountService.GetStatement(c.Request.Context(), id, dateFrom, dateTo, format)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Data: statement,
	})
}
