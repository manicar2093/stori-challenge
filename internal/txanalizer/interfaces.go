package txanalizer

import (
	"github.com/google/uuid"
	"github.com/jordan-wright/email"
)

type (
	TransactionRepository interface {
		Create(input CreateAccountTransactionsInput) error
	}

	Notificator interface {
		SendAccountDetailsEmail(input SendAccountDetailsEmailInput) error
	}

	EmailSender func(*email.Email, MailgunConfig) error

	UUIDCreator func() uuid.UUID
)
