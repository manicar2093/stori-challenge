package txanalizer

import (
	"encoding/csv"

	"github.com/jszwec/csvutil"
	"github.com/manicar2093/filestores"
	"github.com/manicar2093/stori-challenge/pkg/filesrepo"
)

type DefaultService struct {
	transactionsRepo TransactionRepository
	emailSender      Notificator
	filestore        filesrepo.FileStore
	uuidGenerator    UUIDCreator
}

func NewDefaultService(
	transactionsRepo TransactionRepository,
	emailSender Notificator,
	filestor filesrepo.FileStore,
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
) error {
	transactions, objectInfo, err := c.getTransactionsFromCsv(input)
	if err != nil {
		return err
	}

	analizys := doAnalizys(transactions)

	if err := c.transactionsRepo.Create(CreateAccountTransactionsInput{
		Transactions: transactions,
		AccountId:    c.uuidGenerator(),
	}); err != nil {
		return err
	}

	if err := c.emailSender.SendAccountDetailsEmail(SendAccountDetailsEmailInput{
		TransactionsAnalyzis: analizys,
		TransactionsCsvFile:  objectInfo.Reader,
	}); err != nil {
		return err
	}

	return nil
}

func (c *DefaultService) getTransactionsFromCsv(
	input AnalyzeAccountTransactionsInput,
) ([]Transaction, filestores.ObjectInfo, error) {
	objectInfo, err := c.filestore.Get(input.TransactionsFilePath)
	if err != nil {
		return nil, filestores.ObjectInfo{}, err
	}

	decoder, err := csvutil.NewDecoder(csv.NewReader(objectInfo.Reader))
	if err != nil {
		return nil, objectInfo, err
	}
	var transactions []Transaction
	if err := decoder.Decode(&transactions); err != nil {
		return nil, objectInfo, err
	}

	return transactions, objectInfo, nil
}

func doAnalizys(transactions []Transaction) TransactionsAnalizys {
	var (
		analizys     = NewTransactionAnalizys()
		creditAmount float64
		creditCount  uint
		debitAmount  float64
		debitCount   uint
	)
	for _, tx := range transactions { //nolint: varnamelen
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

	return *analizys
}
