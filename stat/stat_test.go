package stat

import (
	"fmt"
	"testing"

	g "github.com/thenorthnate/grace"
)

func TestMax(t *testing.T) {
	z := g.Zeros(g.TypeFloat64, 3, 3, 3)
	mx := Max(z)
	fmt.Println(mx)
}
