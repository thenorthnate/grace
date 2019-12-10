package grace

import "fmt"

// VkF64 implements the float 64 implementation of the grace interface
type VkF64 struct {
	slc []float64
	mat [][]float64
}

// TypeFloat64 returns a grace interface with subtype of float64
func TypeFloat64() Grace {
	return &VkF64{}
}

// Slc returns the slice in the grace interface
func (vv *VkF64) Slc() interface{} {
	return vv.slc
}

// Mat returns the matrix in the grace interface
func (vv *VkF64) Mat() interface{} {
	return vv.mat
}

// GetType returns the data type of the grace structure
func (vv *VkF64) GetType() string {
	slcType := fmt.Sprintf("%T", vv.slc)
	return slcType[2:]
}

// MkSlc initializes the array with 0 values
func (vv *VkF64) MkSlc(shape []int) {
	var r, c int
	if len(shape) == 2 {
		r = shape[0]
		c = shape[1]
	} else if len(shape) == 1 {
		r = 1
		c = shape[0]
	}
	vv.slc = make([]float64, r*c, r*c)
	vv.Reshape(r, c)

}

func (vv *VkF64) Display(depth int) {

}

// Cp makes a copy of the data structure
func (vv *VkF64) Cp() Grace {
	r := len(vv.mat)
	c := len(vv.mat[0])

	dest := make([]float64, len(vv.slc), len(vv.slc))
	_ = copy(dest, vv.slc)
	vv2 := VkF64{
		slc: dest,
	}
	vv2.Reshape(r, c)
	return &vv2
}

// mkF64 creates a new float64 leaf
func mkF64(shape ...int) *VkF64 {
	leaf := VkF64{}
	leaf.MkSlc(shape)
	return &leaf
}

// Reshape shapes the matrix to the given dimensions
func (vv *VkF64) Reshape(r, c int) {
	if r*c != len(vv.slc) {
		return
	}
	vv.mat = make([][]float64, r, r)
	for i := range vv.mat {
		vv.mat[i] = vv.slc[i*c : (i+1)*c]
	}
}
