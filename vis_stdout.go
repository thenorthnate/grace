package grace

// visualize the data to stdout

// Show replaces "head" and "tail" for other frameworks. It prints out the matrix to the depth specified.
// If depth == 0, it will print the entire matrix. If depth is negative, it acts like "tail". Positive
// values will print out the first rows of a matrix.
func (vk *Vektr) Show(depth int) {
	vk.g.Display(depth)
}

/*
// Display prints the matrix to make it visible
func (vk *Vektr) Display(depth int) {
	vkShow(vk, 0, depth)
	if len(vk.shape) == 1 {
		fmt.Println()
	}
}

func vkShow(vk *Vektr, level, depth int) {
	if vk.ptr != nil {
		// has ptr matrices
		for i, v := range vk.ptr {
			if level == 0 {
				if inBounds(i, len(vk.ptr), depth) {
					vkShow(v, level+1, depth)
					fmt.Println()
				}
			} else {
				vkShow(v, level+1, depth)
			}
		}
	} else {
		fmt.Printf("%v", vk.slc)
	}

}

func inBounds(idx int, mLen int, bounds ...int) bool {
	if len(bounds) == 1 {
		// just a simple head or tail operation
		boundary := bounds[0]
		if boundary > 0 {
			if idx < boundary {
				return true
			}
		} else if boundary < 0 {
			if idx >= mLen+boundary {
				return true
			}
		}
	}
	return false
}
*/
