package vkf64

import (
	"fmt"
	"testing"
)

func TestZeros(t *testing.T) {
	v := Zeros(3, 3)
	fmt.Println(v.shape)
	v.Display(-3)
	fmt.Printf("Type: %T\n", v)

	value, err := v.At(0, 0)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Value is: ", value)
}
