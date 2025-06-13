package handlers

import (
	"net/http"

	"finance-api/internal/models"
	"finance-api/internal/services"

	"github.com/gin-gonic/gin"
)

type StatsHandler struct {
	statsService *services.StatsService
}

func NewStatsHandler(statsService *services.StatsService) *StatsHandler {
	return &StatsHandler{
		statsService: statsService,
	}
}

func (h *StatsHandler) GetSummary(c *gin.Context) {
	summary, err := h.statsService.GetSummary(c.Request.Context())
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Data: summary,
	})
}
