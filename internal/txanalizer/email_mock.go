package txanalizer

import (
	"github.com/manicar2093/stori-challenge/pkg/logger"
)

type EmailMock struct{}

func NewEmailMock() *EmailMock {
	return &EmailMock{}
}

func (c EmailMock) SendAccountDetailsEmail(input SendAccountDetailsEmailInput) error {
	tplContent, err := renderEmailTemplate(input.TransactionsAnalyzis)
	if err != nil {
		return err
	}

	logger.GetLogger().WithField("html", string(tplContent)).Info("email sending was mocked")
	return nil
}
