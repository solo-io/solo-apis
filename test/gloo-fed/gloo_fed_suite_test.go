package gloo_fed_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestGlooFed(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gloo Fed Suite")
}
