package grace

import (
	"errors"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	data := []uint8{1, 2, 3, 4, 5}
	arr, err := New(data)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(arr)
}

func TestTrimTo(t *testing.T) {
	data := []uint8{1, 1, 1, 2, 2, 2, 1, 1, 1, 2, 2, 2}
	arr, err := New(data)
	if err != nil {
		t.Error(err.Error())
	}
	err = Reshape(arr, 2, 2, 3)
	if err != nil {
		t.Error(err.Error())
	}
	narr, err := Get(arr, 0, 1)
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(narr)
}

func sampleDoubleApplyFunc(s []int, slc interface{}) (NdArray, error) {
	useSlc, ok := slc.([]uint8)
	if !ok {
		return nil, errors.New("invalid type")
	}
	newSlc := make([]uint8, len(useSlc), len(useSlc))
	for i, v := range useSlc {
		newSlc[i] = v * 2
	}
	return New(newSlc)
}

func TestApply(t *testing.T) {

}
