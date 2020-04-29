package grace

import (
	"errors"
	"fmt"
)

// Ndu8 is a unit8 vekter implementation
type Ndu8 struct {
	s   []int
	slc []uint8
}

func (nd *Ndu8) dtype() string {
	return fmt.Sprintf("%T", nd.slc)
}

func (nd *Ndu8) shape(dims ...int) []int {
	if len(dims) > 0 {
		nd.s = dims
	}
	return nd.s
}

func (nd *Ndu8) setSlc(data interface{}) error {
	switch v := data.(type) {
	case []uint8:
		nd.slc = v
	case int:
		nd.slc = make([]uint8, v, v)
	default:
		return errors.New("invalid data type for slc")
	}
	return nil
}

func (nd *Ndu8) trimTo(start, end int) error {
	if start < 0 || start >= end || end > len(nd.slc) {
		return errors.New("invalid trim region")
	}
	nd.slc = nd.slc[start:end]
	return nil
}

func (nd *Ndu8) get(start, end int) (NdArray, error) {
	if start < 0 || start >= end || end > len(nd.slc) {
		return nil, errors.New("invalid trim region")
	}
	// si := end - start
	// nslc := make([]uint8, si, si)
	// _ = copy(nslc, nd.slc[start:end])
	// n, err := New(nslc)
	n, err := New(nd.slc[start:end])
	return n, err
}

func (nd *Ndu8) arange(start, end, step uint8) {
	size := (end - start) / step
	nd.setSlc(int(size))
	i := 0
	for v := start; v < end; v = v + step {
		nd.slc[i] = v
		i++
	}
}

func (nd *Ndu8) cp() *Ndu8 {
	ssize := len(nd.s)
	dsize := len(nd.slc)
	new := Ndu8{
		s:   make([]int, ssize, ssize),
		slc: make([]uint8, dsize, dsize),
	}

	_ = copy(new.s, nd.s)
	_ = copy(new.slc, nd.slc)
	return &new
}
