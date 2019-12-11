package grace

// Trying yet again

import "errors"

// Narr supports uint8
type Narr struct {
	slc []uint8
	mat [][]uint8
	ptr []*Narr
}

func Empty() *Narr {
	return &Narr{}
}

func Arr(d []uint8) *Narr {
	nda := Narr{
		slc: d,
	}
	return &nda
}

func (nda *Narr) Sub(sarrs ...*Narr) *Narr {
	nda.ptr = make([]*Narr, len(sarrs), len(sarrs))
	for i := range sarrs {
		nda.ptr[i] = Empty()
	}
}

func (g *Gu8) arange(params ...int) error {
	// [start, end, step]
	// [start, end]
	// [end] (0 - end, up by 1)
	var start int
	var end int
	var step int

	l := len(params)
	switch l {
	case 1:
		start = 0
		end = params[0]
		step = 1
	case 2:
		start = params[0]
		end = params[1]
		step = 1
	case 3:
		start = params[0]
		end = params[1]
		step = params[2]
	default:
		return errors.New("invalid params")
	}
	if end > 255 || start > 255 {
		return errors.New("uint8 only supports 0-255")
	} else if end < 0 || start < 0 {
		return errors.New("uint8 can only be positive")
	}
	size := (end - start) / step
	g.slc = make([]uint8, size, size)
	i := 0
	for v := start; v < end; v = v + step {
		g.slc[i] = uint8(v)
		i++
	}
	return nil
}
