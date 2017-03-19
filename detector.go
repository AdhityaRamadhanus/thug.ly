package thugly

import (
	"github.com/lazywei/go-opencv/opencv"
)

// Detector is struct which consist of FaceCascade and EyeCascade to detect face
// region and eyes region of picture, both of them are HaarClasifierCascade
type Detector struct {
	FaceCascade *opencv.HaarCascade
	EyeCascade  *opencv.HaarCascade
}

// NewDetector is Detector constructor, it takes faceCascadePath and eyeCascadePath, path file to your LoadHaarClassifierCascade
func NewDetector(faceCascadePath, eyeCascadePath string) *Detector {
	cascadeFace := opencv.LoadHaarClassifierCascade(faceCascadePath)
	cascadeEye := opencv.LoadHaarClassifierCascade(eyeCascadePath)
	return &Detector{
		FaceCascade: cascadeFace,
		EyeCascade:  cascadeEye,
	}
}

// DetectPersons is function that takes opencv IplImage and returning slice of Persons
func (d *Detector) DetectPersons(image *opencv.IplImage) []Person {
	var persons []Person
	// Detect Faces
	faces := d.FaceCascade.DetectObjects(image)
	// For every faces we determine if there's recognizable eyes in it
	for _, face := range faces {
		// face pointer will be deleted so i create new Rectangle here
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
