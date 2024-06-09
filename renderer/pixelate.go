package renderer

import (
	"image"
	"image/color"
)

// Pixelate pixelates an image with the specified pixel size.
func Pixelate(img image.Image, pixelSize int) image.Image {
	bounds := img.Bounds()
	dst := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y += pixelSize {
		for x := bounds.Min.X; x < bounds.Max.X; x += pixelSize {
			// Average color in the pixel block
			var r, g, b, a uint32
			var count uint32
			for dy := 0; dy < pixelSize && y+dy < bounds.Max.Y; dy++ {
				for dx := 0; dx < pixelSize && x+dx < bounds.Max.X; dx++ {
					rr, gg, bb, aa := img.At(x+dx, y+dy).RGBA()
					r += rr
					g += gg
					b += bb
					a += aa
					count++
				}
			}
			r /= count
			g /= count
			b /= count
			a /= count
			avgColor := color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}

			// Fill the pixel block with the average color
			for dy := 0; dy < pixelSize && y+dy < bounds.Max.Y; dy++ {
				for dx := 0; dx < pixelSize && x+dx < bounds.Max.X; dx++ {
					dst.Set(x+dx, y+dy, avgColor)
				}
			}
		}
	}
	return dst
}
