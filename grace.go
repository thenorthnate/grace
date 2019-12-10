package grace

import (
	"errors"
)

const (
	OpAdd = iota
	OpSub
	OpMul
	OpDiv
	// F64 describes the float64 data type
	F64
	// U8 describes the uint8 data type
	U8
)

// Grace is an interface that supports all operations for matrix manipulations
type Grace interface {
	GetType() string
	Display(depth int)
	MkSlc(shape []int)
}

// Vektr is the grace package equivalent of a vector with ptr-dimensions
type Vektr struct {
	shape []int
	dtype int
	g     Grace
	ptr   []*Vektr
}

// Shape returns the shape of the vektr
func (vk *Vektr) Shape() []int {
	return vk.shape
}

// Size returns the size of the Vektr
func (vk *Vektr) Size() int {
	size := 1
	for _, dim := range vk.shape {
		size *= dim
	}
	return size
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

// DType decodes the datatype into a human readable string
func (vk *Vektr) DType() string {
	switch vk.dtype {
	case F64:
		return "float64"
	default:
		return ""
	}
}

// Vat returns the Vektr at the desired location
func (vk *Vektr) Vat(loc ...int) (*Vektr, error) {
	if len(loc) != len(vk.shape) {
		return vk, errors.New("improper constraints to identify value at location")
	}
	parent := vk
	for _, idx := range loc {
		if parent.ptr != nil {
			if idx <= len(parent.ptr) {
				parent = parent.ptr[idx]
			}
		} else {
			// in the leaf of the tree
			return parent, nil
		}
	}
	return vk, errors.New("location not found")
}

// Reshape the matrix
func (vk *Vektr) Reshape(shape ...int) error {
	return nil
}
