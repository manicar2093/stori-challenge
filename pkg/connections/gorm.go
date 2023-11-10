package connections

import (
	"sync"
	"time"

	"github.com/manicar2093/stori-challenge/pkg/config"
	"github.com/manicar2093/stori-challenge/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

var (
	singleConn *DBWPaginator
	once       sync.Once
)

type DBWPaginator struct {
	*gorm.DB
	PageSize int64
}

func GetGormConnection(pageSize int64, maxIdleConns, maxOpenConns int) *DBWPaginator {
	if singleConn == nil {
		once.Do(func() {
			gormDB, err := gorm.Open(postgres.Open(config.Instance.DatabaseURL))
			if err != nil {
				logger.GetLogger().Panicln(err)
			}
			sqlDB, err := gormDB.DB()
			if err != nil {
				logger.GetLogger().Panicln(err)
			}
			sqlDB.SetMaxIdleConns(maxIdleConns)
			sqlDB.SetMaxOpenConns(maxOpenConns)
			sqlDB.SetConnMaxLifetime(time.Hour)
			singleConn = &DBWPaginator{DB: gormDB, PageSize: pageSize}
		})
	}
	return singleConn
}
