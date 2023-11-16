package txanalizer_test

import (
	"strings"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/jordan-wright/email"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"

	"github.com/manicar2093/stori-challenge/internal/txanalizer"
	"github.com/manicar2093/stori-challenge/mocks"
)

var _ = Describe("Email", func() {

	var (
		expectedEmailTo       string
		expectedEmailFrom     string
		expectedSmtpServer    string
		expectedSmtpAddr      string
		expectedSmtpUser      string
		expectedSmtpPassword  string
		expectedSmtpHost      string
		emailSenderMock       *mocks.EmailSender
		expectedMailgunConfig txanalizer.MailgunConfig
		service               *txanalizer.Mailgun
	)

	BeforeEach(func() {
		expectedEmailTo = "aToEmail@test.com"
		expectedEmailFrom = "aFromEmail@test.com"
		expectedSmtpServer = "smtp.test.com"
		expectedSmtpAddr = "smtp.test.com:587"
		expectedSmtpUser = "smtpUser"
		expectedSmtpPassword = "smtpPassword"
		expectedSmtpHost = "smtpHost"
		emailSenderMock = mocks.NewEmailSender(GinkgoT())
		expectedMailgunConfig = txanalizer.MailgunConfig{
			EmailTo:      expectedEmailTo,
			EmailFrom:    expectedEmailFrom,
			SmtpServer:   expectedSmtpServer,
			SmtpAddr:     expectedSmtpAddr,
			SmtpUser:     expectedSmtpUser,
			SmtpPassword: expectedSmtpPassword,
			SmtpHost:     expectedSmtpHost,
		}
		service = txanalizer.NewMailgun(emailSenderMock.Execute, expectedMailgunConfig)
	})

	Describe("SendAccountDetailsEmail", func() {
		It("sends data to injected email", func() {
			var (
				expectedTransactionCsvFile = strings.NewReader(`Id,Amount,Date
0,10,7/6
1,10,7/6`)
				expectedHtml         = []byte("asdfasdf")
				expectedInputEmailTo = faker.Email()
				input                = txanalizer.SendAccountDetailsEmailInput{
					SendTo: expectedInputEmailTo,
					TransactionsAnalyzis: txanalizer.TransactionsAnalizys{
						TotalBalance: 100,
						TransactionByMonth: txanalizer.TransactionByMonth{
							time.April:  2,
							time.August: 3,
						},
						AverageDebitAmount:  200,
						AverageCreditAmount: 300,
					},
					TransactionsCsvFile: expectedTransactionCsvFile,
				}
			)
			expectedEmail := email.NewEmail()
			expectedEmail.Attach(expectedTransactionCsvFile, "transations.csv", "text/csv")
			expectedEmail.To = []string{expectedEmailTo, expectedInputEmailTo}
			expectedEmail.From = expectedEmailFrom
			expectedEmail.Subject = "Your account status"
			expectedEmail.HTML = expectedHtml
			emailSenderMock.EXPECT().Execute(mock.AnythingOfType("*email.Email"), expectedMailgunConfig).Return(nil)

			Expect(service.SendAccountDetailsEmail(input)).To(Succeed())
		})
	})

})
