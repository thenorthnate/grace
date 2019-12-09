package vkf64

import (
	"errors"
	"fmt"
)

// Vektr is the grace equivalent of a vector with ptr-dimensions
type Vektr struct {
	shape []int
	arr   []float64
	mat   [][]float64
	ptr   []*Vektr
}

// Ptr returns the pointer to the linked vektr
func (vk *Vektr) Ptr() []*Vektr {
	return vk.ptr
}

func (vk *Vektr) PtrInit(shape []int) {

}

func (vk *Vektr) ArrInit(shape []int) {

}

// Shape returns the shape of the vektr
func (vk *Vektr) Shape() []int {
	return vk.shape
}

// Zeros empties the matrix values, and returns a new array of zeros
func Zeros(shape ...int) *Vektr {
	vk := Vektr{
		shape: shape,
	}
	build(&vk, shape...)
	return &vk
}

func build(parent *Vektr, shape ...int) {
	if len(shape) > 1 {
		// not at the last dimension!
		parent.ptr = make([]*Vektr, shape[0], shape[0])
		for i := range parent.ptr {
			parent.ptr[i] = &Vektr{
				shape: shape[1:],
			}
			build(parent.ptr[i], shape[1:]...)
		}
	} else {
		parent.arr = make([]float64, shape[0], shape[0])
	}
}

// IsLeaf returns the truthiness of whether the vektr is a leaf of the data structure
func (vk *Vektr) IsLeaf() bool {
	if vk.ptr == nil {
		return true
	}
	return false
}

// Display prints the matrix to make it visible
func (vk *Vektr) Display(depth int) {
	vkShow(vk, 0, depth)
	if len(vk.shape) == 1 {
		fmt.Println()
	}
}

func vkShow(vk *Vektr, level, depth int) {
	if vk.ptr != nil {
		// has ptr matrices
		for i, v := range vk.ptr {
			if level == 0 {
				if inBounds(i, len(vk.ptr), depth) {
					vkShow(v, level+1, depth)
					fmt.Println()
				}
			} else {
				vkShow(v, level+1, depth)
			}
		}
	} else {
		fmt.Printf("%v", vk.arr)
	}

}

func inBounds(idx int, mLen int, bounds ...int) bool {
	if len(bounds) == 1 {
		// just a simple head or tail operation
		boundary := bounds[0]
		if boundary > 0 {
			if idx < boundary {
				return true
			}
		} else if boundary < 0 {
			if idx >= mLen+boundary {
				return true
			}
		}
	}
	return false
}

// At returns the value at the desired location
func (vk *Vektr) At(loc ...int) (float64, error) {
	if len(loc) != len(vk.shape) {
		return 0, errors.New("improper constraints to identify value at location")
	}
	parent := vk
	for _, idx := range loc {
		if parent.ptr != nil {
			if idx <= len(parent.ptr) {
				parent = parent.ptr[idx]
			}
		} else {
			// in the leaf of the tree
			return parent.arr[idx], nil
		}
	}
	return 0, errors.New("location not found")
}
