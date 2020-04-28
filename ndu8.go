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

func (nd *Ndu8) initSlc(size int) {
	nd.slc = make([]uint8, size, size)
}

func (nd *Ndu8) setSlc(data interface{}) error {
	d, ok := data.([]uint8)
	if !ok {
		return errors.New("invalid data type for slc")
	}
	nd.slc = d
	return nil
}

func (nd *Ndu8) get(loc ...int) error {
	// [
	//	[[1, 1, 1], [2, 2, 2]],
	//  [[1, 1, 1], [2, 2, 2]]
	// ]
	// 2x2x3
	// [1, 1, 1, 2, 2, 2, 1, 1, 1, 2, 2, 2]
	// get(0) --> [[1, 1, 1], [2, 2, 2]]
	// get(0, 0) --> [1, 1, 1]
	// get(0, 0, 0) --> 1
}

func (nd *Ndu8) arange(start, end, step uint8) error {
	size := (end - start) / step
	nd.initSlc(int(size))
	i := 0
	for v := start; v < end; v = v + step {
		nd.slc[i] = v
		i++
	}
}
