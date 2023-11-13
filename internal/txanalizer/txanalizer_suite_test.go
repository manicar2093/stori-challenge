package txanalizer_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTxanalizer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Txanalizer Suite")
}
