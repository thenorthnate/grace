package grace

import (
	"fmt"
	"testing"
)

func TestZeros(t *testing.T) {
	a := &Arru8{}
	Zeros(a, 1, 12)

	err := Reshape(a, 3, 4)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(a)

	fmt.Println(a.Value)
	fmt.Printf("%T, %v, %v\n", a.Value, len(a.Value), cap(a.Value))
	fmt.Printf("%T, %v, %v\n", a.Value[0], len(a.Value[0]), cap(a.Value[0]))

}

func TestPlus(t *testing.T) {
	a := &Arru8{}
	Zeros(a, 1, 12)

	b := &Arru8{}
	Zeros(b, 1, 12)

	_ = Reshape(a, 3, 4)

	err := a.plus(uint8(14))
	if err != nil {
		t.Error(err)
	}
	_ = b.plus(uint8(1))

	err = a.minus(b)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(a.Value)
	fmt.Println(b.Value)
}
