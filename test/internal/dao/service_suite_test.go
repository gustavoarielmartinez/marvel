package dao

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestHandlerSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dao Tests Suite")
}
