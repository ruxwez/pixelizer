package renderer

import (
	"image"
)

// Pixelate pixelates an image with the specified pixel size.
func Pixelate(img image.Image, pixelSize int) image.Image {
	bounds := img.Bounds()
	dst := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y += pixelSize {
		for x := bounds.Min.X; x < bounds.Max.X; x += pixelSize {
			// Get the color of the center pixel in the block
			centerX := x + pixelSize/2
			centerY := y + pixelSize/2
			if centerX >= bounds.Max.X {
				centerX = bounds.Max.X - 1
			}
			if centerY >= bounds.Max.Y {
				centerY = bounds.Max.Y - 1
			}
			blockColor := img.At(centerX, centerY)

			// Fill the pixel block with the center pixel color
			for dy := 0; dy < pixelSize && y+dy < bounds.Max.Y; dy++ {
				for dx := 0; dx < pixelSize && x+dx < bounds.Max.X; dx++ {
					dst.Set(x+dx, y+dy, blockColor)
				}
			}
		}
	}
	return dst
}
