package actor_controller

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestActorController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ActorController Suite")
}
