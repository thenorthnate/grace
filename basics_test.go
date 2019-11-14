package grace

import (
	"fmt"
	"testing"
)

func TestZeros(t *testing.T) {
	a := NewU8([][]uint8{[]uint8{1, 2, 3}, []uint8{3, 2, 1}})
	fmt.Println(a)

	_, b := a.Sum(1)
	fmt.Printf("Sum is: %v (%T)\n", b, b)

	_, c := a.Mean(1)
	fmt.Printf("Mean is: %v (%T)\n", c, c)
}
