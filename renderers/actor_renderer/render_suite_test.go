package actor_renderer

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEntityRenderer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Entity ActorRenderer Suite")
}
