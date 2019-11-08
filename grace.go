package grace

import (
	"errors"
)

// Arr interface exports the common functionality between different array types
type Arr interface {
	zeros(r, c int)
	reshape(r, c int) error
}

// Zeros initializes a matrix of zeros in the specified shape (row, col)
func Zeros(a Arr, r, c int) {
	a.zeros(r, c)
}

// Reshape changes the dimensions of the matrix to the provided values
func Reshape(g Arr, r, c int) error {
	if g == nil {
		return errors.New("array is nil... cannot reshape")
	}
	return g.reshape(r, c)
}
