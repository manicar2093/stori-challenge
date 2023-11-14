package txanalizer_test

import (
	"testing"

	"github.com/manicar2093/stori-challenge/internal/txanalizer"
	"github.com/manicar2093/stori-challenge/pkg/connections"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var conn *connections.DBWPaginator

func TestTxanalizer(t *testing.T) {
	defer GinkgoRecover()
	conn = connections.GetGormConnection()
	if err := conn.AutoMigrate(&txanalizer.Transaction{}); err != nil {
		Fail(err.Error())
	}
	RegisterFailHandler(Fail)
	RunSpecs(t, "Txanalizer Suite")
	if res := conn.Delete(&txanalizer.Transaction{}, "true"); res.Error != nil {
		Fail(res.Error.Error())
	}
}
