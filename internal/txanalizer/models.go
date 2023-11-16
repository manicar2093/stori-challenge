package txanalizer

import (
	"io"
	"time"

	"github.com/google/uuid"
)

type (
	Transaction struct {
		Id        uint      `csv:"-" gorm:"primaryKey;autoIncrement;->"`
		AccountId uuid.UUID `csv:"-"`
		Date      Date      `csv:"Date"`
		Amount    float64   `csv:"Transaction"`
	}

	TransactionByMonth map[time.Month]uint

	TransactionsAnalizys struct {
		TotalBalance        float64
		TransactionByMonth  TransactionByMonth
		AverageDebitAmount  float64
		AverageCreditAmount float64
	}

	CreateAccountTransactionsInput struct {
		Transactions []Transaction
		AccountId    uuid.UUID
	}

	SendAccountDetailsEmailInput struct {
		TransactionsAnalyzis TransactionsAnalizys
		TransactionsCsvFile  io.Reader
		SendTo               string
	}

	AnalyzeAccountTransactionsInput struct {
		TransactionsFilePath string `json:"-"`
		SendTo               string `json:"send_to" validate:"required|email"`
	}

	AnalyzeAccountTransactionsOutput struct {
		EmailId uuid.UUID
	}
)

func NewTransactionAnalizys() *TransactionsAnalizys {
	return &TransactionsAnalizys{
		TransactionByMonth: make(TransactionByMonth),
	}
}
