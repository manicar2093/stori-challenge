package connections

import (
	"database/sql"
	"sync"

	_ "github.com/libsql/libsql-client-go/libsql"
	"github.com/manicar2093/stori-challenge/pkg/config"
	"github.com/manicar2093/stori-challenge/pkg/logger"
	_ "modernc.org/sqlite"
)

var (
	singleConn *sql.DB
	once       sync.Once
)

func GetTursoConnection() *sql.DB {
	if singleConn == nil {
		once.Do(func() {
			db, err := sql.Open("libsql", config.Instance.DatabaseURL)
			if err != nil {
				logger.GetLogger().Panicln(err)
			}
			singleConn = db
		})
	}
	return singleConn
}
