package commands

import (
	"image"
	"image/draw"
	"image/jpeg"
	"log"
	"os"

	thugly "github.com/AdhityaRamadhanus/thug.ly"
	"github.com/disintegration/imaging"
	"github.com/lazywei/go-opencv/opencv"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var (
	FaceCascadePath = "cascades/haarcascade_frontalface_alt.xml"
	EyeCascadePath  = "cascades/haarcascade_eye.xml"
	GlassesPath     = "assets/glasses.png"
)

func fatalErr(err error) {
	log.Println(err)
	os.Exit(1)
}

func Thuglify(cliContext *cli.Context) {
	imagePath := cliContext.String("input")
	if len(imagePath) == 0 {
		fatalErr(errors.New("Please provide input"))
	}
	detector := thugly.NewDetector(FaceCascadePath, EyeCascadePath)
	baseImage, err := thugly.LoadImage(imagePath)
	if err != nil {
		fatalErr(errors.Wrap(err, "Failed to load your image"))
	}
	persons := detector.DetectPersons(opencv.FromImage(baseImage))
	glassesImage, err := thugly.LoadImage(GlassesPath)
	if err != nil {
		fatalErr(errors.Wrap(err, "Failed to load the glsses image"))
	}
	bounds := baseImage.Bounds()
	canvas := image.NewRGBA(bounds)
	draw.Draw(canvas, bounds, baseImage, bounds.Min, draw.Over)
	for _, person := range persons {
		eyesReg := person.GetEyesRect()
		fitGlassImg := imaging.Resize(glassesImage, eyesReg.Dx(), eyesReg.Dy(), imaging.Lanczos)
		draw.Draw(
			canvas,
			eyesReg,
			fitGlassImg,
			bounds.Min,
			draw.Over)
	}
	outputImage, err := os.Create(cliContext.String("output"))
	if err != nil {
		fatalErr(errors.Wrap(err, "Cannot open output image"))
	}
	if err := jpeg.Encode(outputImage, canvas, &jpeg.Options{jpeg.DefaultQuality}); err != nil {
		fatalErr(errors.Wrap(err, "Cannot encode jpeg output"))
	}

}
