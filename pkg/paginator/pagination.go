package paginator

import (
	"errors"
	"math"

	"gorm.io/gorm"
)

const (
	pageSizeUpperLimit = 100
	pageSizeLowerLimit = 10
)

var (
	ErrPageNotExists = errors.New("requested page cannot be created")
)

type Page[T any] struct {
	CurrentPage int64 `json:"current_page"`
	PageSize    int64 `json:"page_size"`
	// TotalEntries indicates how many data exists
	TotalEntries int64 `json:"total_entries"`
	TotalPages   int64 `json:"total_pages"`
	NextPage     int64 `json:"next_page"`
	// EntriesCount indicates how many data the current page contains
	EntriesCount int64 `json:"entries_count"`
	Data         []T   `json:"data"`
}

func (c *Page[T]) SelectPages(query *gorm.DB) error {
	var model T
	query.Model(&model).Count(&c.TotalEntries)
	c.checkPageSizeLimits()
	c.setTotalPages()
	if err := c.checkPageCanBeCreated(); err != nil {
		return err
	}
	if res := query.Model(&model).Scopes(func(db *gorm.DB) *gorm.DB {
		offset := c.getOffset()
		return db.Offset(offset).Limit(int(c.PageSize))
	}).Find(&c.Data); res.Error != nil {
		return res.Error
	}
	c.calculateNextPage()
	c.setEntriesCount()
	return nil
}

func (c *Page[T]) SelectPagesRow(countRawQuery, rawQuery *gorm.DB) error {
	countRawQuery.Scan(&c.TotalEntries)
	c.checkPageSizeLimits()
	c.setTotalPages()
	if err := c.checkPageCanBeCreated(); err != nil {
		return err
	}

	if res := rawQuery.Offset(c.getOffset()).Limit(int(c.PageSize)).Scan(&c.Data); res.Error != nil {
		return res.Error
	}
	c.calculateNextPage()
	c.setEntriesCount()
	return nil
}

func (c *Page[T]) HasNextPage() bool {
	return c.CurrentPage < c.TotalPages
}

func (c *Page[T]) setTotalPages() {
	var (
		totalEntriesF = float64(c.TotalEntries)
		pageSizeF     = float64(c.PageSize)
		totalPagesF   = math.Ceil(totalEntriesF / pageSizeF)
	)
	c.TotalPages = int64(totalPagesF)
	if c.TotalPages == 0 {
		c.TotalPages = 1
	}
}

func (c *Page[T]) checkPageSizeLimits() {
	if c.PageSize > pageSizeUpperLimit {
		c.PageSize = pageSizeUpperLimit
		return
	}
	if c.PageSize < pageSizeLowerLimit {
		c.PageSize = pageSizeLowerLimit
		return
	}
}

func (c *Page[T]) checkPageCanBeCreated() error {
	if c.CurrentPage > c.TotalPages {
		return ErrPageNotExists
	}
	return nil
}

func (c *Page[T]) calculateNextPage() {
	if len(c.Data) == 0 {
		c.NextPage = 1
		return
	}
	if c.CurrentPage == c.TotalPages {
		c.NextPage = 1
		return
	}
	c.NextPage = c.CurrentPage + 1
}

func (c *Page[T]) setEntriesCount() {
	c.EntriesCount = int64(len(c.Data))
}

func (c *Page[T]) getOffset() int {
	return int((c.CurrentPage - 1) * c.PageSize)
}
