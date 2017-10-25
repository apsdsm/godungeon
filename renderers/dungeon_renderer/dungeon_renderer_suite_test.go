package dungeon_renderer_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDungeonRenderer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DungeonRenderer Suite")
}
