package ai_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AI Suite")
}