package grace

import (
	"fmt"
	"testing"
)

func TestVektr(t *testing.T) {
	z := Zeros(TypeFloat64, 3, 3, 3)
	vk, err := AsF64(z.ptr[0].g)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(vk.mat)

	fmt.Println(z.DType())

}

func TestGu8(t *testing.T) {
	g := Gu8{}
	g.arange(0, 10, 2)
	fmt.Println(g.slc)

	g.arange(10)
	fmt.Println(g.slc)

	g.arange(1, 5)
	fmt.Println(g.slc)

	a := Empty()
}
