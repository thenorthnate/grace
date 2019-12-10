package grace

import (
	"errors"
)

// AsF64 returns the typed grace component as a float64 struct
func AsF64(g Grace) (*VkF64, error) {
	vk, ok := g.(*VkF64)
	if !ok {
		return vk, errors.New("invalid type check... must convert explicitly")
	}
	return vk, nil
}
