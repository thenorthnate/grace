package grace

import (
	"errors"
)

// NdArray is a generic interface for ndarray type operations
type NdArray interface {
	dtype() string
	shape(dims ...int) []int
	initSlc(size int)
	setSlc(data interface{}) error
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
	vv.setShape(dims...)
}

// Zeros creates a new matrix of zeros with the given dimensions and data type
func Zeros(nda NdArray, dims ...int) {
	s := dimSize(dims...)
	nda.mkSlc(s)
	nda.setShape(dims...)
}

// New creates a new NdArray with the given data
func New(data interface{}) (NdArray, error) {
	var nda NdArray
	switch v := data.(type) {
	case []uint8:
		nda = &Ndu8{}
	default:
		return nil, errors.New("unsupported input data type")
	}
	err := nda.setSlc(data)
	return nda, err
}

func Arange(params ...interface{}) (NdArray, error) {
	ap, err := mkArangeParams(params)
	if err != nil {
		return nil, err
	}
	f64slc, err := arangeF64(ap...)
	if err != nil {
		return nil, err
	}

	var nda NdArray
	switch v := params.(type) {
	case []uint8:
		nda = &Ndu8{}
	default:
		return nil, errors.New("unsupported input data type")
	}
	return nda, nil
}

func arangeF64(params ...float64) ([]float64, error) {
	if len(params) != 3 {
		return []float64{}, errors.New("requires 3 input parameters")
	}
	start := params[0]
	end := params[1]
	step := params[2]
	size := (end - start) / step
	slc = make([]float64, size, size)
	i := 0
	for v := start; v < end; v = v + step {
		slc[i] = v
		i++
	}
	return slc, nil
}

func mkArangeParams(params ...interface{}) ([]float64, error) {
	p := float64{}
	for _, v := range params {
		pv := float64(v)
		p = append(p, pv)
	}
	switch len(p) {
	case 0:
		return []float64{}, errors.New("must provide parameters for arange")
	case 1:
		p = append([]float64{0}, p[0], 1)
	case 2:
		p = append(p, 1)
	case 3:
		// do nothing... all parameters given
	default:
		return []float64{}, errors.New("cannot give more than 3 parameters to arange")
	}
	return p, nil
}
