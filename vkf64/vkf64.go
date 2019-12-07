package vkf64

import (
	"errors"
	"fmt"
)

// Vektr is the grace equivalent of a vector with sub-dimensions
type Vektr struct {
	v     []float64
	sub   []*Vektr
	shape []int
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
		parent.sub = make([]*Vektr, shape[0], shape[0])
		for i := range parent.sub {
			parent.sub[i] = &Vektr{
				shape: shape[1:],
			}
			build(parent.sub[i], shape[1:]...)
		}
	} else {
		parent.v = make([]float64, shape[0], shape[0])
	}
}

// IsLeaf returns the truthiness of whether the vektr is a leaf of the data structure
func (vk *Vektr) IsLeaf() bool {
	if vk.sub == nil {
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
	if vk.sub != nil {
		// has sub matrices
		for i, v := range vk.sub {
			if level == 0 {
				if inBounds(i, len(vk.sub), depth) {
					vkShow(v, level+1, depth)
					fmt.Println()
				}
			} else {
				vkShow(v, level+1, depth)
			}
		}
	} else {
		fmt.Printf("%v", vk.v)
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
		if parent.sub != nil {
			if idx <= len(parent.sub) {
				parent = parent.sub[idx]
			}
		} else {
			// in the leaf of the tree
			return parent.v[idx], nil
		}
	}
	return 0, errors.New("location not found")
}
