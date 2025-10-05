package canvas_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCanvas(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Canvas Suite")
}
