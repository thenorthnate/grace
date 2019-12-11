package grace

import (
	"errors"
)

const (
	// OpAdd defines an add operation
	OpAdd = iota
	// OpSub defines an subtraction operation
	OpSub
	// OpMul defines an multiply operation
	OpMul
	// OpDiv defines an divide operation
	OpDiv
)

// Grace is an interface that supports all operations for matrix manipulations
type Grace interface {
	GetType() string
	Display(depth int)
	MkSlc(shape []int)
	Cp() Grace
	Slc() interface{}
	Mat() interface{}
}

// Ops defines the base statistical operations
type Ops interface {
	Max() interface{}
}

// Vektr is the grace package equivalent of a vector with ptr-dimensions
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

// G returns the grace sub interface
func (vk *Vektr) G() Grace {
	return vk.g
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
	return vk.dtype
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

// Execute applies a recursive execution of the given function to leaf nodes
func Execute(vk *Vektr, execF func(inVk *Vektr, params interface{}), execParams interface{}) {
	if !vk.IsLeaf() {
		for i := range vk.ptr {
			Execute(vk.ptr[i], execF, execParams)
		}
	} else {
		// is a leaf
		execF(vk, execParams)
	}

}
