package thugly

import (
	"image"
	"image/color"
	"os"
)

func LoadImage(file string) (image.Image, error) {
	reader, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	return img, nil
}

// For Debugging purpose, you mostly won't need this

func drawHLine(image *image.RGBA, x1, y, x2 int, isBlack bool) {
	col := color.White
	if isBlack {
		col = color.Black
	}
	for ; x1 <= x2; x1++ {
		image.Set(x1, y, col)
	}
}

func drawVLine(image *image.RGBA, y1, x, y2 int, isBlack bool) {
	col := color.White
	if isBlack {
		col = color.Black
	}
	for ; y1 <= y2; y1++ {
		image.Set(x, y1, col)
	}
}

func drawRect(image *image.RGBA, x1, y1, x2, y2 int, isBlack bool) {
	drawHLine(image, x1, y1, x2, isBlack)
	drawHLine(image, x1, y2, x2, isBlack)
	drawVLine(image, y1, x1, y2, isBlack)
	drawVLine(image, y1, x2, y2, isBlack)
}
