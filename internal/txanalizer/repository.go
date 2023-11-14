package txanalizer

import (
	"github.com/huandu/go-sqlbuilder"
	"github.com/manicar2093/stori-challenge/pkg/connections"
)

type TursoRepository struct {
	conn *connections.DBWPaginator
}

func NewTursoRepository(conn *connections.DBWPaginator) *TursoRepository {
	return &TursoRepository{
		conn: conn,
	}
}

func (c *TursoRepository) Create(input CreateAccountTransactionsInput) error {
	builder := sqlbuilder.SQLite.NewInsertBuilder()
	builder.InsertInto("transactions")
	builder.Cols("id", "amount", "date", "account_id")
	for _, data := range input.Transactions {
		builder.Values(data.Id, data.Amount, data.Date, input.AccountId)
	}

	sql, args := builder.Build()

	if res := c.conn.Exec(sql, args...); res.Error != nil {
		return res.Error
	}

	return nil
}
