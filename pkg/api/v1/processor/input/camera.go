package input

import (
	"errors"
	"gocv.io/x/gocv"
)

type Camera struct {
	cam   *gocv.VideoCapture
	frame gocv.Mat
}

func (o *Camera) Init(params map[string]interface{}) (err error) {
	o.cam, err = gocv.VideoCaptureDevice(0)
	if err != nil {
		return err
	}

	o.frame = gocv.NewMat()

	return
}

func (o *Camera) Close() (err error) {
	err = o.cam.Close()
	if err != nil {
		return err
	}

	err = o.frame.Close()
	if err != nil {
		return err
	}

	return
}

func (o *Camera) Read() (frame *gocv.Mat, err error) {
	if !o.cam.Read(&o.frame) {
		err = errors.New("cannot read from cam")
		return nil, err
	}

	frame = &o.frame

	return
}

func (o *Camera) Process(frame *gocv.Mat) (err error) {
	return
}
