package txanalizer

import (
	"bytes"
	"html/template"
	"net/smtp"
	"time"

	"github.com/coditory/go-errors"
	"github.com/jordan-wright/email"
)

var monthsInSpanish = map[time.Month]string{
	time.January:   "Enero",
	time.February:  "Febrero",
	time.March:     "Marzo",
	time.April:     "Abril",
	time.May:       "Mayo",
	time.June:      "Junio",
	time.July:      "Julio",
	time.August:    "Agosto",
	time.September: "Septiembre",
	time.October:   "Octubre",
	time.November:  "Noviembre",
	time.December:  "Diciembre",
}

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
	tplContent, err := renderEmailTemplate(input.TransactionsAnalyzis)
	if err != nil {
		return errors.Wrap(err)
	}

	newEmail := email.NewEmail()
	newEmail.Attach(input.TransactionsCsvFile, "transations.csv", "text/csv")
	newEmail.To = []string{c.config.EmailTo}
	newEmail.From = c.config.EmailFrom
	newEmail.Subject = "Your account status"
	newEmail.HTML = tplContent
	return c.sender(newEmail, c.config)
}

func renderEmailTemplate(transactionAnalizys TransactionsAnalizys) ([]byte, error) {
	tpl, err := template.New("emailContent").Funcs(template.FuncMap{
		"getMonthName": func(month time.Month) string {
			return monthsInSpanish[month]
		},
	}).Parse(accountStatusHtmlTemplate)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	tplContent := new(bytes.Buffer)
	if err := tpl.Execute(tplContent, transactionAnalizys); err != nil {
		return nil, errors.Wrap(err)
	}

	return tplContent.Bytes(), nil
}
