package output

import (
	"gocv.io/x/gocv"
	"errors"
)

type IMShow struct {
	Label string `json:"label"`

	window *gocv.Window
}

func (o *IMShow) Default() (err error) {
	o.Label= "Result"

	return
}

func (o *IMShow) Init(params map[string]interface{}) (err error) {
	if label, ok := params["label"]; ok {
		o.Label, ok = label.(string)
		if !ok {
			err = errors.New("label is not string type")
			return err
		}
	}

	o.window = gocv.NewWindow(o.Label)

	return
}

func (o *IMShow) Close() (err error) {
	err = o.window.Close()
	if err != nil {
		return err
	}

	return
}

func (o *IMShow) Process(frame *gocv.Mat) (err error) {
	o.window.IMShow(*frame)
	if o.window.WaitKey(1) >= 0 {
		err = errors.New("window was closed")
		return err
	}

	return
}
