package handlers

import (
	"finance-api/internal/services"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	AccountHandler  *AccountHandler
	TransferHandler *TransferHandler
	StatsHandler    *StatsHandler
}

func NewHandlers(
	accountService *services.AccountService,
	transferService *services.TransferService,
	statsService *services.StatsService,
) *Handlers {
	return &Handlers{
		AccountHandler:  NewAccountHandler(accountService),
		TransferHandler: NewTransferHandler(transferService),
		StatsHandler:    NewStatsHandler(statsService),
	}
}

func SetupRouter(handlers *Handlers) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(LoggingMiddleware())
	router.Use(ErrorHandlerMiddleware())
	router.Use(CORSMiddleware())
	router.Use(RateLimitMiddleware())

	v1 := router.Group("/api/v1")
	{
		accounts := v1.Group("/accounts")
		{
			accounts.POST("", handlers.AccountHandler.CreateAccount)
			accounts.GET("", handlers.AccountHandler.GetAccounts)
			accounts.GET("/:id", handlers.AccountHandler.GetAccount)
			accounts.GET("/:id/balance", handlers.AccountHandler.GetBalance)
			accounts.POST("/:id/deposit", handlers.AccountHandler.Deposit)
			accounts.POST("/:id/withdraw", handlers.AccountHandler.Withdraw)
			accounts.GET("/:id/statement", handlers.AccountHandler.GetStatement)
		}

		transfers := v1.Group("/transfers")
		{
			transfers.POST("", handlers.TransferHandler.CreateTransfer)
			transfers.GET("", handlers.TransferHandler.GetTransfers)
			transfers.GET("/:id", handlers.TransferHandler.GetTransfer)
		}

		stats := v1.Group("/stats")
		{
			stats.GET("/summary", handlers.StatsHandler.GetSummary)
		}
	}

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Service is running",
		})
	})

	return router
}
