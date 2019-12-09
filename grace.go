package grace

import (
	"errors"
)

const (
	OpAdd = iota
	OpSub
	OpMul
	OpDiv

	U8  = "uint8"
	F64 = "float64"
	F32 = "float32"
)

// Grace is an interface that supports all operations for matrix manipulations
type Grace interface {
	Display(depth int)
	MkSlc(shape []int)
}

// Vektr is the grace equivalent of a vector with ptr-dimensions
type Vektr struct {
	shape []int
	dtype string
	g     Grace
	ptr   []*Vektr
}

// Shape returns the shape of the vektr
func (vk *Vektr) Shape() []int {
	return vk.shape
}

// Ptr returns the pointer to the linked vektr
func (vk *Vektr) Ptr() []*Vektr {
	return vk.ptr
}

// IsLeaf returns the truthiness of whether the vektr is a leaf of the data structure
func (vk *Vektr) IsLeaf() bool {
	if vk.ptr == nil {
		return true
	}
	return false
}

// Show replaces "head" and "tail" for other frameworks. It prints out the matrix to the depth specified.
// If depth == 0, it will print the entire matrix. If depth is negative, it acts like "tail". Positive
// values will print out the first rows of a matrix.
func Show(g Grace, depth int) {
	g.Display(depth)
}

// At returns the value at the desired location
func At(g Grace, loc ...int) (float64, error) {
	if len(loc) != len(g.Shape()) {
		return 0, errors.New("improper constraints to identify value at location")
	}
	parent := g
	for _, idx := range loc {
		tmpPs := parent.Ptr()
		if tmpPs != nil {
			if idx <= len(tmpPs) {
				parent = tmpPs[idx]
			}
		} else {
			// in the leaf of the tree
			return parent.v[idx], nil
		}
	}
	return 0, errors.New("location not found")
}
