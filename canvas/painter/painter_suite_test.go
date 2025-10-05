package painter_test

import (
	"bytes"

	"github.com/apsdsm/canvas"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"

	"testing"
)

func TestPainter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Painter Suite")
}

// get a string from a layer
func getLayerLine(layer *canvas.Layer, x, y, len int) string {

	var buffer bytes.Buffer

	for r := 0; r < len; r++ {
		c := layer.Grid[x+r][y].Rune

		// if the rune isn't set to the null value
		if c != 0 {
			_, _ = buffer.WriteRune(c)
		}
	}

	return buffer.String()
}

// print the code points in a string
func debugString(input string) {
	fmt.Printf("% x\n", input)
}
