package thugly

import (
	"image"
	"image/color"
	"io/ioutil"
	"os"

	"github.com/golang/freetype"
)

// LoadImage is a helper to load image file and returning image.Image and error
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

func DrawLabel(img *image.RGBA, label string, font string, rect image.Rectangle) error {
	// fixed for now
	dpi := 72.0
	// Calculate fontsize
	fontsize := float64(rect.Dx()/len(label)) * 2.52
	// Read the font data.
	fontBytes, err := ioutil.ReadFile(font)
	if err != nil {
		return err
	}
	fontType, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return err
	}

	fontColor := image.Black
	c := freetype.NewContext()
	c.SetDPI(dpi)
	c.SetFont(fontType)
	c.SetFontSize(fontsize)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(fontColor)

	fontPoint := freetype.Pt(rect.Min.X, rect.Min.Y)
	_, err = c.DrawString(label, fontPoint)
	return err
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
