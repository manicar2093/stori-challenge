package txanalizer_test

import (
	"os"
	"time"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/manicar2093/filestores"
	"github.com/manicar2093/stori-challenge/internal/txanalizer"
	"github.com/manicar2093/stori-challenge/mocks"
)

var _ = Describe("DefaultService", func() {

	var (
		transactionsRepoMock *mocks.TransactionRepository
		emailSenderMock      *mocks.EmailSender
		fileRepoMock         *mocks.FileStore
		uuidGeneratorMock    *mocks.UUIDCreator
		service              *txanalizer.DefaultService
	)

	BeforeEach(func() {
		T := GinkgoT()
		transactionsRepoMock = mocks.NewTransactionRepository(T)
		emailSenderMock = mocks.NewEmailSender(T)
		uuidGeneratorMock = mocks.NewUUIDCreator(T)
		fileRepoMock = mocks.NewFileStore(T)
		service = txanalizer.NewDefaultService(transactionsRepoMock, emailSenderMock, fileRepoMock, uuidGeneratorMock.Execute)
	})

	Describe("AnalyzeAccountTransactions", func() {
		It("sends an email with calculated account data", func() {
			var (
				expectedTransactionsFilePath = "account_details_test.csv"
				expectedTransactionCsvFile   *os.File
				expectedObjectInfo           filestores.ObjectInfo
				transactions                 = []txanalizer.Transaction{
					{
						Id:     0,
						Date:   txanalizer.NewDate(2023, time.July, 15, 0, 0, 0, 0, time.UTC),
						Amount: 60.5,
					},
					{
						Id:     1,
						Date:   txanalizer.NewDate(2023, time.July, 28, 0, 0, 0, 0, time.UTC),
						Amount: -10.3,
					},
					{
						Id:     2,
						Date:   txanalizer.NewDate(2023, time.August, 2, 0, 0, 0, 0, time.UTC),
						Amount: -20.46,
					},
					{
						Id:     3,
						Date:   txanalizer.NewDate(2023, time.August, 13, 0, 0, 0, 0, time.UTC),
						Amount: 10,
					},
				}
				expectedUuidReturn             = uuid.New()
				createAccountTransactionsInput = txanalizer.CreateAccountTransactionsInput{
					Transactions: transactions,
					AccountId:    expectedUuidReturn,
				}
				sendAccountDetailsEmailInput = txanalizer.SendAccountDetailsEmailInput{
					TransactionsAnalyzis: txanalizer.TransactionsAnalizys{
						TotalBalance: 39.74,
						TransactionByMonth: txanalizer.TransactionByMonth{
							time.July:   2,
							time.August: 2,
						},
						AverageDebitAmount:  -15.38,
						AverageCreditAmount: 35.25,
					},
					TransactionsCsvFile: expectedTransactionCsvFile,
				}
				emailIdReturn = uuid.New()
				input         = txanalizer.AnalyzeAccountTransactionsInput{
					TransactionsFilePath: expectedTransactionsFilePath,
				}
			)
			fileRepoMock.EXPECT().Get(expectedTransactionsFilePath).Return(expectedObjectInfo, nil)
			transactionsRepoMock.EXPECT().Create(&createAccountTransactionsInput).Return(nil)
			emailSenderMock.EXPECT().SendAccountDetailsEmail(sendAccountDetailsEmailInput).Return(emailIdReturn, nil)
			uuidGeneratorMock.EXPECT().Execute().Return(emailIdReturn)

			got, err := service.AnalyzeAccountTransactions(input)

			Expect(err).ToNot(HaveOccurred())
			Expect(got.EmailId).To(Equal(emailIdReturn))
		})
	})

})
