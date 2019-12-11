package grace

// Max returns the maximum value in the grace interface
func (vv *VkU8) Max() interface{} {
	var mv uint8
	for _, v := range vv.slc {
		if v > mv {
			mv = v
		}
	}
	return mv
}
