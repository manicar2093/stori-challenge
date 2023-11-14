package txanalizer_test

import (
	"time"

	"github.com/google/uuid"
	"github.com/manicar2093/stori-challenge/internal/txanalizer"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Turso", func() {

	var (
		repo *txanalizer.TursoRepository
	)

	BeforeEach(func() {
		repo = txanalizer.NewTursoRepository(conn)
	})

	Describe("Create", func() {
		It("saves all account transactions", func() {
			var (
				expectedAccountId = uuid.New()
				input             = txanalizer.CreateAccountTransactionsInput{
					Transactions: []txanalizer.Transaction{
						{Id: 0, AccountId: expectedAccountId, Date: txanalizer.NewDate(time.November, 11)},
						{Id: 1, AccountId: expectedAccountId, Date: txanalizer.NewDate(time.November, 12)},
						{Id: 2, AccountId: expectedAccountId, Date: txanalizer.NewDate(time.November, 13)},
					},
					AccountId: expectedAccountId,
				}
			)

			Expect(repo.Create(input)).To(Succeed())

			var found txanalizer.Transaction
			if err := conn.QueryRow(
				"SELECT * FROM TRANSACTIONS WHERE id = 0",
			).Scan(&found.Id, &found.Amount, &found.Date, &found.AccountId); err != nil {
				Fail(err.Error())
			}
			Expect(found.AccountId).To(Equal(expectedAccountId.String()))
		})
	})

})
