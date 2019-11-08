package grace

import (
	"errors"
	"fmt"
)

// Arru8 is the uint8 Arr implementation
type Arru8 struct {
	Data  []uint8
	Value [][]uint8
}

func (a *Arru8) reshape(r, c int) error {
	if r*c != len(a.Data) {
		return errors.New("invalid shape specified")
	}
	a.Value = make([][]uint8, r, r)
	for i := range a.Value {
		a.Value[i] = a.Data[i*c : (i+1)*c]
	}
	return nil
}

func (a *Arru8) zeros(r, c int) {
	a.Data = make([]uint8, r*c, r*c)
	a.reshape(r, c)
}

func (a *Arru8) mul(b *Arru8) error {
	if len(a.Data) != len(b.Data) {
		return errors.New("arrays must be the same length")
	}
	for i, val := range b.Data {
		a.Data[i] *= val
	}
	return nil
}

func (a *Arru8) div(b *Arru8) error {
	if len(a.Data) != len(b.Data) {
		return errors.New("arrays must be the same length")
	}
	for i, val := range b.Data {
		// Should i catch div/0 errors or let them cause a panic?
		a.Data[i] /= val
	}
	return nil
}

func (a *Arru8) plus(b interface{}) error {
	switch v := b.(type) {
	case uint8:
		for i := range a.Data {
			a.Data[i] += v
		}
	case *Arru8:
		if len(a.Data) != len(v.Data) {
			return errors.New("arrays must be the same length")
		}
		for i, val := range v.Data {
			a.Data[i] += val
		}
	default:
		return fmt.Errorf("input not compatible with base array (type: %T)", a.Data[0])
	}
	return nil
}

func (a *Arru8) minus(b interface{}) error {
	switch v := b.(type) {
	case uint8:
		for i := range a.Data {
			a.Data[i] -= v
		}
	case *Arru8:
		if len(a.Data) != len(v.Data) {
			return errors.New("arrays must be the same length")
		}
		for i, val := range v.Data {
			a.Data[i] -= val
		}
	default:
		return fmt.Errorf("input not compatible with base array (type: %T)", a.Data[0])
	}
	return nil
}
