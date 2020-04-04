package api_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestConnectTheSpan(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ConnectTheSpan Suite")
}
