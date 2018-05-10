package transform

import (
	"gocv.io/x/gocv"
	"image"
	"errors"
)

type Resizer struct {
	Scale float64
}

func (o *Resizer) Init(params map[string]interface{}) (err error) {
	scale := params["scale"]
	switch scale.(type) {
	case int:
		o.Scale = float64(scale.(int))
	case float64:
		o.Scale = scale.(float64)
	default:
		return errors.New("scale is neither int nor float")
	}

	return
}

func (o *Resizer) Process(frame *gocv.Mat) (err error) {
	if o.Scale != 1 {
		gocv.Resize(*frame, frame, image.Point{}, o.Scale, o.Scale, gocv.InterpolationLinear)
	}

	return
}
