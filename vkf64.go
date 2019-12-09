package grace

// VKf64 implements the float 64 implementation of the grace interface
type VKf64 struct {
	slc []float64
	mat [][]float64
}

// Slc returns the pointer to the linked vektr
func (vk *VKf64) Slc() []float64 {
	return vk.slc
}

// Mat returns the pointer to the linked vektr
func (vk *VKf64) Mat() [][]float64 {
	return vk.mat
}

// MkSlc initializes the array with 0 values
func (vk *VKf64) MkSlc(shape []int) {
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
func (vk *VKf64) Reshape(r, c int) {
	if r*c != len(vk.slc) {
		return
	}
	vk.mat = make([][]float64, r, r)
	for i := range vk.mat {
		vk.mat[i] = vk.slc[i*c : (i+1)*c]
	}
}

func (vk *VKf64) Display(depth int) {

}

// ToF64 returns the typed grace component as a float64 struct
func ToF64(g Grace) *VKf64 {
	vk, _ := g.(*VKf64)
	return vk
}
