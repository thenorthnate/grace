package pic

import (
	"image"

	"github.com/thenorthnate/grace"
)

const (
	channelsNRGBA = 4
)

func FromNRGBA(img *image.NRGBA) *[]grace.Arru8 {
	out := make([]grace.Arru8{}, channelsNRGBA, channelsNRGBA)
	ccl := len(img.Pix) / channelsNRGBA // color channel length
	for i := 0; i < channelsNRGBA; i++ {
		out[i] = make(grace.Arru8, ccl, ccl)
	}

	for i := range out[0] {
		for j := 0; j < channelsNRGBA; j++ {
			out[j].Data[i] = img.Pix[(i+j)*channelsNRGBA]
		}
	}

	r := len(img.Pix) / img.Stride
	c := img.Stride / channelsNRGBA
	for i := 0; i < channelsNRGBA; i++ {
		grace.Reshape(&out[i], r, c)
	}
	return &out
}
