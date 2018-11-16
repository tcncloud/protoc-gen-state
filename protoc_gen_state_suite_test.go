package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestProtocGenState(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProtocGenState Suite")
}
