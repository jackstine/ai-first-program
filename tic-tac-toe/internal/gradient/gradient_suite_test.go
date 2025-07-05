package gradient_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGradient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gradient Suite")
}