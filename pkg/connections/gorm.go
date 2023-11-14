package connections

import (
	"log"
	"sync"

	libsql "github.com/ekristen/gorm-libsql"
	"gorm.io/gorm"

	"github.com/manicar2093/stori-challenge/pkg/config"
	_ "modernc.org/sqlite"
)

type DBWPaginator struct {
	*gorm.DB
	PageSize int64
}

var (
	singleGormConn *gorm.DB
	onceGorm       sync.Once
)

func GetGormConnection() *DBWPaginator {
	if singleGormConn == nil {
		onceGorm.Do(func() {
			gormDB, err := gorm.Open(libsql.Open(config.Instance.DatabaseURL))
			if err != nil {
				log.Panicln(err)
			}
			singleGormConn = gormDB
		})
	}
	return &DBWPaginator{DB: singleGormConn}

}
