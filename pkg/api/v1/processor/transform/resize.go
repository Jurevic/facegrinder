package transform

import (
	"errors"
	"gocv.io/x/gocv"
	"image"
)

type Resizer struct {
	Scale float64 `json:"scale"`
}

func (o *Resizer) Default() (err error) {
	o.Scale = 1

	return
}

func (o *Resizer) Init(params map[string]interface{}) (err error) {
	if scale, ok := params["scale"]; ok {
		o.Scale, ok = scale.(float64)
		if !ok {
			err = errors.New("scale is not float type")
			return err
		}
	}

	return
}

func (o *Resizer) Process(frame *gocv.Mat) (err error) {
	if o.Scale != 1 {
		gocv.Resize(*frame, frame, image.Point{}, o.Scale, o.Scale, gocv.InterpolationLinear)
	}

	return
}
