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

func mustOrFail[T any](ret T, err error) T {
	if err != nil {
		Fail(err.Error())
	}
	return ret
}

var _ = Describe("DefaultService", func() {

	var (
		transactionsRepoMock *mocks.TransactionRepository
		notificatorMock      *mocks.Notificator
		fileRepoMock         *mocks.FileStore
		uuidGeneratorMock    *mocks.UUIDCreator
		service              *txanalizer.DefaultService
	)

	BeforeEach(func() {
		T := GinkgoT()
		transactionsRepoMock = mocks.NewTransactionRepository(T)
		notificatorMock = mocks.NewNotificator(T)
		uuidGeneratorMock = mocks.NewUUIDCreator(T)
		fileRepoMock = mocks.NewFileStore(T)
		service = txanalizer.NewDefaultService(transactionsRepoMock, notificatorMock, fileRepoMock, uuidGeneratorMock.Execute)
	})

	Describe("AnalyzeAccountTransactions", func() {
		It("sends an email with calculated account data", func() {
			var (
				expectedTransactionsFilePath = "account_details_test.csv"
				expectedTransactionCsvFile   = mustOrFail(os.Open(expectedTransactionsFilePath))
				expectedObjectInfo           = mustOrFail(filestores.FileToStoreInfo(expectedTransactionCsvFile))
				transactions                 = []txanalizer.Transaction{
					{
						Id:     0,
						Date:   txanalizer.NewDate(time.July, 15),
						Amount: 60.5,
					},
					{
						Id:     1,
						Date:   txanalizer.NewDate(time.July, 28),
						Amount: -10.3,
					},
					{
						Id:     2,
						Date:   txanalizer.NewDate(time.August, 2),
						Amount: -20.46,
					},
					{
						Id:     3,
						Date:   txanalizer.NewDate(time.August, 13),
						Amount: 10,
					},
				}
				expectedAccountUuid            = uuid.New()
				createAccountTransactionsInput = txanalizer.CreateAccountTransactionsInput{
					Transactions: transactions,
					AccountId:    expectedAccountUuid,
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
					TransactionsCsvFile: expectedObjectInfo.Reader,
				}
				input = txanalizer.AnalyzeAccountTransactionsInput{
					TransactionsFilePath: expectedTransactionsFilePath,
				}
			)
			fileRepoMock.EXPECT().Get(expectedTransactionsFilePath).Return(expectedObjectInfo, nil)
			transactionsRepoMock.EXPECT().Create(createAccountTransactionsInput).Return(nil)
			notificatorMock.EXPECT().SendAccountDetailsEmail(sendAccountDetailsEmailInput).Return(nil)
			uuidGeneratorMock.EXPECT().Execute().Return(expectedAccountUuid)

			err := service.AnalyzeAccountTransactions(input)

			Expect(err).ToNot(HaveOccurred())
		})
	})

})
