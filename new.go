package grace

import (
	"fmt"
)

// Zeros creates a new matrix of zeros with the given dimensions and data type
func Zeros(dtype int, shape ...int) *Vektr {
	vk := Vektr{shape: shape}
	build(&vk, dtype, shape...)
	return &vk
}

// build is a recurrsive function that initializes a grace interface
func build(parent *Vektr, dtype int, shape ...int) {
	parent.dtype = dtype
	if len(shape) > 2 {
		// not at the last 2 dimensions!
		parent.mkPtr(shape) // creates structures and sets shape
		for i := range parent.ptr {
			build(parent.ptr[i], dtype, shape[1:]...)
		}
	} else {
		// At the final one or two dimensions
		parent.mkLeaf()
	}
}

// mkPtr initializes a new slice of sub nodes
func (vk *Vektr) mkPtr(shape []int) {
	vk.ptr = make([]*Vektr, shape[0], shape[0])
	for i := range vk.ptr {
		vk.ptr[i] = &Vektr{
			shape: shape[1:],
		}
	}
}

func (vk *Vektr) mkLeaf() {
	switch vk.dtype {
	case F64:
		vk.g = mkF64(vk.shape...)
	default:
		fmt.Println("unrecognized data type")
	}
}

// Duplicate copies all the values of the Vektr to a new variable in memory
// and returns a new pointer to that structure
func (vk *Vektr) Duplicate() *Vektr {
	return vk
}
