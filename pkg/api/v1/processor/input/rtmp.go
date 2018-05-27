package input

import (
	"errors"
	"gocv.io/x/gocv"
)

type Rtmp struct {
	Url string `json:"url"`

	stream *gocv.VideoCapture
	frame  gocv.Mat
}

func (o *Rtmp) Default() (err error) {
	o.Url = "rtmp://localhost/1"

	return
}
func (o *Rtmp) Init(params map[string]interface{}) (err error) {
	if url, ok := params["url"]; ok {
		o.Url, ok = url.(string)
		if !ok {
			err = errors.New("url is not string type")
			return err
		}
	}

	o.stream, err = gocv.VideoCaptureFile(o.Url)
	if err != nil {
		err = errors.New("could not initialise specified channel")
		return err
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

func (o *Rtmp) Process(frame *gocv.Mat) (err error) {
	return
}
