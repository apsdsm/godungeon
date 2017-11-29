package actor_renderer_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEntityRenderer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ActorRenderer Suite")
}
