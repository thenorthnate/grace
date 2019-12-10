package grace

// Provides functionality to get the 5 number summary of the data set

func (vk *Vektr) Max() float64 {
	return 0
}

func (vk *Vektr) Min() float64 {
	return 0
}

func (vk *Vektr) Mean() float64 {
	return 0
}

func (vk *Vektr) Median() float64 {
	return 0
}

func (vk *Vektr) UpperQ() float64 {
	return 0
}

func (vk *Vektr) LowerQ() float64 {
	return 0
}

/*
import (
	"errors"
	"fmt"
)


// Get returns a slice to the underlying matrix. If provided, it will return just the indicies specified.
func (a *MatrixU8) Get(loc ...int) [][]uint8 {
	for _, v := range loc {
		if v < 0 {
			fmt.Println("warning: locations must be positive values")
			return a.m
		}
	}
	if len(loc) == 2 {
		// given r and c
		return [][]uint8{
			[]uint8{a.m[loc[0]][loc[1]]},
		}
	} else if len(loc) == 4 {
		// r1, r2, c1, c2
		rowC := loc[1] - loc[0]
		out := make([][]uint8, rowC, rowC)
		for i := range out {
			out[i] = a.m[i+loc[0]][loc[2]:loc[3]]
		}
		return out
	}
	return a.m
}

// Set places the given value at the row and column specified
func (a *MatrixU8) Set(r, c int, v uint8) {
	a.m[r][c] = v
}

// Reshape shapes the matrix to the given dimensions
func (a *MatrixU8) Reshape(r, c int) error {
	if r*c != len(a.v) {
		return errors.New("invalid shape specified")
	}
	a.m = make([][]uint8, r, r)
	for i := range a.m {
		a.m[i] = a.v[i*c : (i+1)*c]
	}
	return nil
}

func (a *MatrixU8) algebraV(op int, v uint8) {
	switch op {
	case OpAdd:
		for i := range a.v {
			a.v[i] += v
		}
	case OpSub:
		for i := range a.v {
			a.v[i] -= v
		}
	case OpMul:
		for i := range a.v {
			a.v[i] *= v
		}
	case OpDiv:
		for i := range a.v {
			a.v[i] /= v
		}
	}
}

func (a *MatrixU8) algebraM(op int, v *MatrixU8) error {
	if len(a.v) != len(v.v) {
		return errors.New("arrays must be the same length")
	}
	switch op {
	case OpAdd:
		for i, val := range v.v {
			a.v[i] += val
		}
	case OpSub:
		for i, val := range v.v {
			a.v[i] -= val
		}
	case OpMul:
		for i, val := range v.v {
			a.v[i] *= val
		}
	case OpDiv:
		for i, val := range v.v {
			a.v[i] /= val
		}
	}
	return nil
}

// Algebra performs element-wise algebra for given values to the "a" matrix.
// Options for "op" include OpAdd, OpSub, OpMul, and OpDiv
// To keep the math fast, there is limited error handling (it is up to the
// user to ensure there is no divide by 0 or overflow, etc)
func (a *MatrixU8) Algebra(b interface{}, op int) error {
	switch v := b.(type) {
	case uint8:
		a.algebraV(op, v)
	case *MatrixU8:
		err := a.algebraM(op, v)
		return err
	default:
		return fmt.Errorf("input not compatible with base array (type: %T)", a.v[0])
	}
	return nil
}

// Sum returns the total sum of the matrix (axis=0 is rows, axis=1 is columns, axis=2 is scalar)
func (a *MatrixU8) Sum(axis int) (int64, []int64) {
	var s int64
	var sr []int64

	switch axis {
	case 0:
		// sum rows
		sr = make([]int64, len(a.m), len(a.m))
		for i := range a.m {
			for _, v := range a.m[i] {
				sr[i] += int64(v)
			}
		}
	case 1:
		// sum columns
		sr = make([]int64, len(a.m[0]), len(a.m[0]))
		for i := range a.m {
			for j, v := range a.m[i] {
				sr[j] += int64(v)
			}
		}
	case 2:
		s = 0
		for _, v := range a.v {
			s += int64(v)
		}
	}

	return s, sr
}

// Mean returns the average value of the matrix
func (a *MatrixU8) Mean(axis int) (float64, []float64) {
	var m float64
	var lv int
	var mv []float64

	s, sr := a.Sum(axis)
	switch axis {
	case 0:
		// rows
		lv = len(a.m[0])
	case 1:
		// cols
		lv = len(a.m)
	case 2:
		m = float64(s) / float64(len(a.v))
	}

	srl := len(sr)
	mv = make([]float64, srl, srl)
	for i, v := range sr {
		mv[i] = float64(v) / float64(lv)
	}

	return m, mv
}

// Dot performs an in-place dot product between matrix A and B
func (a *MatrixU8) Dot(b *MatrixU8) error {
	return nil
}

// Identity turns "a" into an identity matrix
func (a *MatrixU8) Identity() {

}
*/
