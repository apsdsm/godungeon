package math2d_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMath2d(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Math2d Suite")
}
