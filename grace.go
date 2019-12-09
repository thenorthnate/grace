package grace

import (
	"errors"
)

const (
	OpAdd = iota
	OpSub
	OpMul
	OpDiv

	gtU8  = "uint8"
	gtF64 = "float64"
	gtF32 = "float32"
)

// Grace is an interface that supports all operations for matrix manipulations
type Grace interface {
	Display(depth int)
	IsLeaf() bool
	Shape() []int
	Ptr() []Grace
	PtrInit(shape []int)
	SlcInit(shape []int)
	Zeros(shape ...int) Grace
}

// Vektr is the grace equivalent of a vector with ptr-dimensions
type Vektr struct {
	shape []int
	g     Grace
	ptr   []*Vektr
}

// Zeros creates a new matrix of zeros with the given dimensions and data type
func Zeros(dtype string, dims ...int) Grace {
	switch dtype {
	case gtF64:
		return
	}
	return nil
}

// Show replaces "head" and "tail" for other frameworks. It prints out the matrix to the depth specified.
// If depth == 0, it will print the entire matrix. If depth is negative, it acts like "tail". Positive
// values will print out the first rows of a matrix.
func Show(g Grace, depth int) {
	g.Display(depth)
}

// build is a recurrsive function that initializes a grace interface
func build(parent Grace, shape ...int) {
	if len(shape) > 1 {
		// not at the last dimension!
		parent.PtrInit(shape) // should make and create the structures in memory
		ptrs := parent.Ptr()
		for i := range ptrs {
			build(ptrs[i], shape[1:]...)
		}
	} else {
		parent.SlcInit(shape)
	}
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
