package thugly

import (
	"image"

	"github.com/lazywei/go-opencv/opencv"
)

type Person struct {
	Face opencv.Rect
	Eyes []opencv.Rect
}

func (p *Person) GetEyesRect() image.Rectangle {
	eyesRegion := p.getDefaultEyesRect()
	leftEyes := []opencv.Rect{}
	rightEyes := []opencv.Rect{}
	// Search Best Candidates for left eye
	for _, leftEyeCandidate := range p.Eyes {
		if (leftEyeCandidate.X() <= p.Face.X()+p.Face.Width()/2) &&
			(leftEyeCandidate.Y() <= p.Face.Y()+p.Face.Height()/2) {
			leftEyes = append(leftEyes, leftEyeCandidate)
		}
	}
	// Search Best Candidates for left eye
	for _, rightEyeCandidate := range p.Eyes {
		if (rightEyeCandidate.X() >= p.Face.X()+p.Face.Width()/2) &&
			(rightEyeCandidate.Y() <= p.Face.Y()+p.Face.Height()/2) {
			rightEyes = append(rightEyes, rightEyeCandidate)
		}
	}

	if len(leftEyes) == 1 && len(rightEyes) == 1 {
		var x1, y1, x2, y2 int
		x1 = leftEyes[0].X() - leftEyes[0].Width()/10
		y1 = leftEyes[0].Y() - leftEyes[0].Height()/10
		x2 = rightEyes[0].X() + rightEyes[0].Width() + rightEyes[0].Width()/10
		y2 = rightEyes[0].Y() + rightEyes[0].Height() + rightEyes[0].Height()/10
		eyesRegion = image.Rect(x1, y1, x2, y2)
	}
	return eyesRegion
}

// Using some magic heuristic here
func (p *Person) getDefaultEyesRect() image.Rectangle {
	var x, y, w, h int
	x = p.Face.X() + (p.Face.Width()*15)/100
	y = p.Face.Y() + (p.Face.Height()*20)/100
	w = p.Face.Width()/2 + (p.Face.Width()*20)/100
	h = p.Face.Height()/2 - (p.Face.Height()*14)/100
	return image.Rect(x, y, x+w, y+h)
}
