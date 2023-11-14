package txanalizer

import (
	"net/smtp"

	"github.com/jordan-wright/email"
)

type (
	Mailgun struct {
		config MailgunConfig
		sender EmailSender
	}

	MailgunConfig struct {
		EmailTo      string
		EmailFrom    string
		SmtpServer   string
		SmtpAddr     string
		SmtpUser     string
		SmtpPassword string
		SmtpHost     string
	}
)

func DefaultSender(e *email.Email, conf MailgunConfig) error {
	return e.Send(
		conf.SmtpAddr,
		smtp.PlainAuth("", conf.SmtpUser, conf.SmtpPassword, conf.SmtpHost),
	)
}

func NewMailgun(sender EmailSender, config MailgunConfig) *Mailgun {
	return &Mailgun{
		config: config,
		sender: sender,
	}
}

func (c Mailgun) SendAccountDetailsEmail(input SendAccountDetailsEmailInput) error {
	newEmail := email.NewEmail()
	newEmail.Attach(input.TransactionsCsvFile, "transations.csv", "text/csv")
	newEmail.To = []string{c.config.EmailTo}
	newEmail.From = c.config.EmailFrom
	newEmail.Subject = "Your account status"
	newEmail.HTML = []byte("asdfasdf")
	return c.sender(newEmail, c.config)
}
