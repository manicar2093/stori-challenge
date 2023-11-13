package txanalizer

import "github.com/google/uuid"

type (
	TransactionRepository interface {
		Create(input CreateAccountTransactionsInput) error
	}

	EmailSender interface {
		SendAccountDetailsEmail(input SendAccountDetailsEmailInput) (uuid.UUID, error)
	}

	UUIDCreator func() uuid.UUID
)
