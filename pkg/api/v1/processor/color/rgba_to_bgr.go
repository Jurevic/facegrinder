package color

import (
	"gocv.io/x/gocv"
)

type RGBAToBGR struct{}

func (o *RGBAToBGR) Process(frame *gocv.Mat) (err error) {
	gocv.CvtColor(*frame, frame, gocv.ColorRGBAToBGR)

	return
}
