package thugly

import (
	"github.com/lazywei/go-opencv/opencv"
)

type Detector struct {
	FaceCascade *opencv.HaarCascade
	EyeCascade  *opencv.HaarCascade
}

func NewDetector(faceCascadePath, eyeCascadePath string) *Detector {
	cascadeFace := opencv.LoadHaarClassifierCascade(faceCascadePath)
	cascadeEye := opencv.LoadHaarClassifierCascade(eyeCascadePath)
	return &Detector{
		FaceCascade: cascadeFace,
		EyeCascade:  cascadeEye,
	}
}

func (d *Detector) DetectPersons(image *opencv.IplImage) []Person {
	var persons []Person
	faces := d.FaceCascade.DetectObjects(image)
	for _, face := range faces {
		roiReg := opencv.NewRect(
			face.X(),
			face.Y(),
			face.Width(),
			face.Height(),
		)
		person := Person{
			Face: roiReg,
		}
		var eyeRect opencv.Rect
		// Detect Eyes
		image.SetROI(roiReg)
		eyes := d.EyeCascade.DetectObjects(image)
		for _, eye := range eyes {
			eyeRect = opencv.NewRect(
				eye.X()+roiReg.X(),
				eye.Y()+roiReg.Y(),
				eye.Width(),
				eye.Height(),
			)
			person.Eyes = append(person.Eyes, eyeRect)
		}
		image.ResetROI()
		persons = append(persons, person)
	}
	return persons
}
