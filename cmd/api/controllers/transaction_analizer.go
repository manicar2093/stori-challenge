package controllers

import (
	"net/http"

	"github.com/coditory/go-errors"
	"github.com/labstack/echo/v4"
	"github.com/manicar2093/stori-challenge/internal/txanalizer"
	"github.com/manicar2093/stori-challenge/pkg/config"
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

//	@Summary		Analyze a transactions file
//	@Description	Analyze a transactions file and send an email with generated data
//	@Tags			transaction_analyzer
//	@Success		200	{object}	echo.Map	"Confirmation message"
//	@Failure		500	"Something unidentified has occurred"
//	@Router			/analyze [post]
func (c *TransactionAnalyzer) PostAnalyzeTransactions(ctx echo.Context) error {
	if err := c.service.AnalyzeAccountTransactions(txanalizer.AnalyzeAccountTransactionsInput{
		TransactionsFilePath: config.Instance.TransactionsFilePath,
	}); err != nil {
		return errors.Wrap(err)
	}
	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "analysis done, email sent",
	})
}
