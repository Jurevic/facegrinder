package input

import (
	"errors"
	"gocv.io/x/gocv"
)

type Rtmp struct {
	stream *gocv.VideoCapture
	frame gocv.Mat
}

func (o *Rtmp) Init(params map[string]interface{}) (err error) {
	url := params["url"].(string)

	o.stream, err = gocv.VideoCaptureFile(url)
	if err != nil {
		return nil
	}

	o.frame = gocv.NewMat()

	return
}

func (o *Rtmp) Close() (err error) {
	err = o.stream.Close()
	if err != nil {
		return err
	}

	err = o.frame.Close()
	if err != nil {
		return err
	}

	return
}

func (o *Rtmp) Read() (frame *gocv.Mat, err error) {
	if !o.stream.Read(&o.frame) {
		err = errors.New("cannot read from rtmp stream")
		return nil, err
	}

	frame = &o.frame

	return
}
