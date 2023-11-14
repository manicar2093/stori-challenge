package txanalizer_test

import (
	"database/sql"
	"testing"

	"github.com/manicar2093/stori-challenge/pkg/connections"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var conn *sql.DB

func TestTxanalizer(t *testing.T) {
	conn = connections.GetTursoConnection()
	RegisterFailHandler(Fail)
	RunSpecs(t, "Txanalizer Suite")
}
