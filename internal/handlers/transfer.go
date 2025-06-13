package handlers

import (
	"net/http"

	"finance-api/internal/models"
	"finance-api/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransferHandler struct {
	transferService *services.TransferService
}

func NewTransferHandler(transferService *services.TransferService) *TransferHandler {
	return &TransferHandler{
		transferService: transferService,
	}
}

func (h *TransferHandler) CreateTransfer(c *gin.Context) {
	var req models.CreateTransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Invalid request",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	transfer, err := h.transferService.CreateTransfer(c.Request.Context(), &req)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusCreated, models.SuccessResponse{
		Data:    transfer,
		Message: "Transfer created successfully",
	})
}

func (h *TransferHandler) GetTransfer(c *gin.Context) {
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

	transfer, err := h.transferService.GetTransfer(c.Request.Context(), id)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Data: transfer,
	})
}

func (h *TransferHandler) GetTransfers(c *gin.Context) {
	var req models.TransfersListRequest
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

	response, err := h.transferService.GetTransfers(c.Request.Context(), &req)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
