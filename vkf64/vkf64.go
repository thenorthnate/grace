package vkf64

import (
	"errors"
	"fmt"
)

// Vektr is the grace equivalent of a vector with ptr-dimensions
type Vektr struct {
	shape []int
	slc   []float64
	mat   [][]float64
	ptr   []*Vektr
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
