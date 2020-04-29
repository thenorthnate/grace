package grace

import (
	"errors"
)

// NdArray is a generic interface for ndarray type operations
type NdArray interface {
	dtype() string
	shape(dims ...int) []int
	setSlc(data interface{}) error
	trimTo(start, end int) error
	get(start, end int) (NdArray, error)
	apply(fn func(s []int, slc interface{}) (NdArray, error)) (NdArray, error)
}

// Shape returns the shape of the NdArray
func Shape(nda NdArray) []int {
	return nda.shape()
}

// Size returns the size of the NdArray
func Size(nda NdArray) int {
	return dimSize(nda.shape()...)
}

// dimSize returns the calculated size from the given dimensions
func dimSize(dims ...int) int {
	size := 1
	for _, d := range dims {
		size *= d
	}
	return size
}

// Dtype decodes the datatype into a human readable string
func Dtype(nda NdArray) string {
	return nda.dtype()
}

// Reshape shapes the matrix to the given dimensions
func Reshape(nda NdArray, dims ...int) error {
	newsize := dimSize(dims...)
	if newsize != Size(nda) {
		return errors.New("invalid shape dimensions given")
	}
	_ = nda.shape(dims...)
	return nil
}

// Zeros creates a new matrix of zeros with the given dimensions and data type
func Zeros(nda NdArray, dims ...int) {
	s := dimSize(dims...)
	_ = nda.setSlc(s)
	_ = nda.shape(dims...)
}

// New creates a new NdArray with the given data
func New(data interface{}) (NdArray, error) {
	var nda NdArray
	switch v := data.(type) {
	case []uint8:
		nda = &Ndu8{
			s:   []int{len(v)},
			slc: v,
		}
	default:
		return nil, errors.New("unsupported input data type")
	}
	// err := nda.setSlc(data)
	return nda, nil
}

// Get returns a new NdArray sliced to the given location
func Get(nda NdArray, loc ...int) (NdArray, error) {
	// new copy of nda2
	// then err := nda2.get(loc...)
	// return nda2, err
	// [
	//	[[1, 1, 1], [2, 2, 2]],
	//  [[1, 1, 1], [2, 2, 2]]
	// ]
	// 2x2x3
	// [1, 1, 1, 2, 2, 2, 1, 1, 1, 2, 2, 2]
	// get(0) --> [[1, 1, 1], [2, 2, 2]] --> 0-6
	// get(0, 0) --> [1, 1, 1] --> 0-3
	// get(0, 0, 0) --> 1 --> 0
	si := Size(nda)
	sh := nda.shape()
	ptr := 0
	for i, v := range loc {
		si /= sh[i]
		ptr += v * si
	}
	// _ = nda.trimTo(ptr, ptr+si)
	return nda.get(ptr, ptr+si)
}

/*
	Loc(nda NdArray, loc ...interface{}) (NdArray, error)
	nda, _ = Loc(nda, "0:4", "10:")
	nda, _ = Loc(nda, 0, 1)
*/
