package grace

import (
	"fmt"
)

// VkU8 implements the uint8 implementation of the grace interface
type VkU8 struct {
	slc []uint8
	mat [][]uint8
}

// TypeUInt8 returns a grace interface with subtype of float64
func TypeUInt8() Grace {
	return &VkU8{}
}

// Slc returns the slice in the grace interface
func (vv *VkU8) Slc() interface{} {
	return vv.slc
}

// Mat returns the matrix in the grace interface
func (vv *VkU8) Mat() interface{} {
	return vv.mat
}

// GetType returns the data type of the grace structure
func (vv *VkU8) GetType() string {
	slcType := fmt.Sprintf("%T", vv.slc)
	return slcType[2:]
}

// MkSlc initializes the array with 0 values
func (vv *VkU8) MkSlc(shape []int) {
	var r, c int
	if len(shape) == 2 {
		r = shape[0]
		c = shape[1]
	} else if len(shape) == 1 {
		r = 1
		c = shape[0]
	}
	vv.slc = make([]uint8, r*c, r*c)
	vv.Reshape(r, c)

}

// Display writes the data to stdout
func (vv *VkU8) Display(depth int) {

}

// Cp makes a copy of the data structure
func (vv *VkU8) Cp() Grace {
	r := len(vv.mat)
	c := len(vv.mat[0])

	dest := make([]uint8, len(vv.slc), len(vv.slc))
	_ = copy(dest, vv.slc)
	vv2 := VkU8{
		slc: dest,
	}
	vv2.Reshape(r, c)
	return &vv2
}

// mkF64 creates a new float64 leaf
func mkU8(shape ...int) *VkU8 {
	leaf := VkU8{}
	leaf.MkSlc(shape)
	return &leaf
}

// Reshape shapes the matrix to the given dimensions
func (vv *VkU8) Reshape(r, c int) {
	if r*c != len(vv.slc) {
		return
	}
	vv.mat = make([][]uint8, r, r)
	for i := range vv.mat {
		vv.mat[i] = vv.slc[i*c : (i+1)*c]
	}
}
