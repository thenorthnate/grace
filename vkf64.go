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

// GetType returns the data type of the grace structure
func (vk *VkF64) GetType() string {
	slcType := fmt.Sprintf("%T", vk.slc)
	return slcType[2:]
}

// mkF64 creates a new float64 leaf
func mkF64(shape ...int) *VkF64 {
	leaf := VkF64{}
	leaf.MkSlc(shape)
	return &leaf
}

// MkSlc initializes the array with 0 values
func (vk *VkF64) MkSlc(shape []int) {
	var r, c int
	if len(shape) == 2 {
		r = shape[0]
		c = shape[1]
	} else if len(shape) == 1 {
		r = 1
		c = shape[0]
	}
	vk.slc = make([]float64, r*c, r*c)
	vk.Reshape(r, c)

}

// Reshape shapes the matrix to the given dimensions
func (vk *VkF64) Reshape(r, c int) {
	if r*c != len(vk.slc) {
		return
	}
	vk.mat = make([][]float64, r, r)
	for i := range vk.mat {
		vk.mat[i] = vk.slc[i*c : (i+1)*c]
	}
}

func (vk *VkF64) Display(depth int) {

}
