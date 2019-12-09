package vkf64

import (
	"errors"
	"fmt"

	"github.com/thenorthnate/grace"
)

// Vektr is the grace equivalent of a vector with ptr-dimensions
type Vektr struct {
	shape []int
	slc   []float64
	mat   [][]float64
	ptr   []*Vektr
}

type VKf64 struct {
	slc []float64
	mat [][]float64
}

// PtrInit initializes a new slice of sub nodes
func (vk *Vektr) PtrInit(shape []int) {
	vk.shape = shape
	vk.ptr = make([]*Vektr, shape[0], shape[0])
	for i := range vk.ptr {
		vk.ptr[i] = &Vektr{
			shape: shape[1:],
		}
	}
}

// SlcInit initializes the array with 0 values
func (vk *Vektr) SlcInit(shape []int) {
	vk.slc = make([]float64, shape[0], shape[0])
}

// Shape returns the shape of the vektr
func (vk *Vektr) Shape() []int {
	return vk.shape
}

// Slc returns the pointer to the linked vektr
func (vk *Vektr) Slc() []float64 {
	return vk.slc
}

// Mat returns the pointer to the linked vektr
func (vk *Vektr) Mat() [][]float64 {
	return vk.mat
}

// Ptr returns the pointer to the linked vektr
func (vk *Vektr) Ptr() []grace.Grace {
	return vk.ptr
}

// Zeros empties the matrix values, and returns a new array of zeros
func Zeros(shape ...int) grace.Grace {
	vk := Vektr{}
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
		parent.slc = make([]float64, shape[0], shape[0])
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
		fmt.Printf("%v", vk.slc)
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
			return parent.slc[idx], nil
		}
	}
	return 0, errors.New("location not found")
}
