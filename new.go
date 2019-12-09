package grace

// Zeros creates a new matrix of zeros with the given dimensions and data type
func Zeros(dtype string, shape ...int) *Vektr {
	vk := Vektr{}
	build(&vk, dtype, shape...)
	return &vk
}

// build is a recurrsive function that initializes a grace interface
func build(parent *Vektr, dtype string, shape ...int) {
	if len(shape) > 2 {
		// not at the last 2 dimensions!
		parent.mkPtr(shape) // creates structures and sets shape
		parent.dtype = dtype
		for i := range parent.ptr {
			build(parent.ptr[i], dtype, shape[1:]...)
		}
	} else {
		// At the final one or two dimensions
		parent.mkLeaf()
	}
}

// mkPtr initializes a new slice of sub nodes
func (vk *Vektr) mkPtr(shape []int) {
	vk.shape = shape
	vk.ptr = make([]*Vektr, shape[0], shape[0])
	for i := range vk.ptr {
		vk.ptr[i] = &Vektr{
			shape: shape[1:],
		}
	}
}

func (vk *Vektr) mkLeaf() {
	switch vk.dtype {
	case F64:
		leaf := VKf64{}
		leaf.MkSlc(vk.shape)
		vk.g = &leaf
	}
}

// Duplicate copies all the values of the Vektr to a new variable in memory
// and returns a new pointer to that structure
func (vk *Vektr) Duplicate() *Vektr {

}
