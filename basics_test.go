package grace

import (
	"fmt"
	"testing"
)

func TestVektr(t *testing.T) {
	z := Zeros(F64, 3, 3, 3)
	vk, err := AsF64(z.ptr[0].g)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(vk.mat)
}
