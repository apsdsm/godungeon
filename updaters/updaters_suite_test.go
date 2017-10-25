package updaters

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestUpdaters(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Updater Suite")
}
