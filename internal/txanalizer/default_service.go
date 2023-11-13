package txanalizer

import filestores "github.com/manicar2093/stori-challenge/pkg/filesrepo"

type DefaultService struct {
	transactionsRepo TransactionRepository
	emailSender      EmailSender
	filestor         filestores.FileStore
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
		filestor:         filestor,
		uuidGenerator:    uuidGenerator,
	}
}

func (c *DefaultService) AnalyzeAccountTransactions(input AnalyzeAccountTransactionsInput) (AnalyzeAccountTransactionsOutput, error) {
	panic("not implemented!")
}
