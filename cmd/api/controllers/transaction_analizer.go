package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/manicar2093/stori-challenge/internal/txanalizer"
)

type TransactionAnalyzer struct {
	service *txanalizer.DefaultService
}

func NewTransactionAnalyzer(service *txanalizer.DefaultService) *TransactionAnalyzer {
	return &TransactionAnalyzer{
		service: service,
	}
}

func (c *TransactionAnalyzer) SetUpRoutes(group *echo.Group) {
	group.POST("/analyze", c.PostAnalyzeTransactions)
}

// @Summary		Gets a json welcome message
// @Description	Return a json content with a welcome message
// @Tags			transaction_analyzer
// @Accept			json
// @Produce		json
// @Success		200	{string}	echo.Map	"Demo data"
// @Failure		500	"Something unidentified has occurred"
// @Router			/ [get]
func (c *TransactionAnalyzer) PostAnalyzeTransactions(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "welcome to the starter template!",
	})
}
