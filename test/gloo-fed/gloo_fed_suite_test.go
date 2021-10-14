package gloo_fed_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGlooFed(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gloo Fed Suite")
}
