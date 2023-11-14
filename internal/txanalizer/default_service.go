package txanalizer

import (
	"encoding/csv"
	"log"

	"github.com/jszwec/csvutil"
	filestores "github.com/manicar2093/stori-challenge/pkg/filesrepo"
)

type DefaultService struct {
	transactionsRepo TransactionRepository
	emailSender      EmailSender
	filestore        filestores.FileStore
	uuidGenerator    UUIDCreator
}

func NewDefaultService(
	transactionsRepo TransactionRepository,
	emailSender EmailSender,
	filestor filestores.FileStore,
	uuidGenerator UUIDCreator,
) *DefaultService {
	return &DefaultService{
		transactionsRepo: transactionsRepo,
		emailSender:      emailSender,
		filestore:        filestor,
		uuidGenerator:    uuidGenerator,
	}
}

func (c *DefaultService) AnalyzeAccountTransactions(
	input AnalyzeAccountTransactionsInput,
) (AnalyzeAccountTransactionsOutput, error) {
	objectInfo, err := c.filestore.Get(input.TransactionsFilePath)
	if err != nil {
		return AnalyzeAccountTransactionsOutput{}, err
	}

	decoder, err := csvutil.NewDecoder(csv.NewReader(objectInfo.Reader))
	if err != nil {
		return AnalyzeAccountTransactionsOutput{}, err
	}
	var transactions []Transaction
	if err := decoder.Decode(&transactions); err != nil {
		return AnalyzeAccountTransactionsOutput{}, err
	}
	var (
		analizys             = NewTransactionAnalizys()
		creditAmount float64 = 0
		creditCount          = 0
		debitAmount  float64 = 0
		debitCount           = 0
	)
	for _, tx := range transactions { //nolint: varnamelen
		log.Println(tx.Date)
		if tx.Amount > 0 {
			creditCount++
			creditAmount += tx.Amount
		} else {
			debitCount++
			debitAmount += tx.Amount
		}
		analizys.TransactionByMonth[tx.Date.Month()]++
		analizys.TotalBalance += tx.Amount
	}
	analizys.AverageCreditAmount = creditAmount / float64(creditCount)
	analizys.AverageDebitAmount = debitAmount / float64(debitCount)

	if err := c.transactionsRepo.Create(CreateAccountTransactionsInput{
		Transactions: transactions,
		AccountId:    c.uuidGenerator(),
	}); err != nil {
		return AnalyzeAccountTransactionsOutput{}, err
	}

	emailId, err := c.emailSender.SendAccountDetailsEmail(SendAccountDetailsEmailInput{
		TransactionsAnalyzis: *analizys,
		TransactionsCsvFile:  objectInfo.Reader,
	})
	if err != nil {
		return AnalyzeAccountTransactionsOutput{}, err
	}

	return AnalyzeAccountTransactionsOutput{
		EmailId: emailId,
	}, nil
}
