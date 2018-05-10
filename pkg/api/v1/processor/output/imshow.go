package output

import (
	"gocv.io/x/gocv"
	"errors"
)

type IMShow struct {
	window *gocv.Window
}

func (o *IMShow) Init(params map[string]interface{}) (err error) {
	name := params["name"].(string)

	o.window = gocv.NewWindow(name)

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
