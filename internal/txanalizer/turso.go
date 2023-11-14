package txanalizer

import (
	"database/sql"
	"errors"

	"github.com/huandu/go-sqlbuilder"
)

type TursoRepository struct {
	conn *sql.DB
}

func NewTursoRepository(conn *sql.DB) *TursoRepository {
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
	println(sql)

	res, err := c.conn.Exec(sql, args...)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected <= 0 {
		return errors.New("no data inserted/any row affected")
	}
	return nil
}
