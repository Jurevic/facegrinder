package context

import (
	"gocv.io/x/gocv"
)

type FrameCpyToCtx struct {
	Key string

	frame gocv.Mat
	context *map[string]interface{}
}

func (o *FrameCpyToCtx) Init(params map[string]interface{}) (err error) {
	o.Key = params["key"].(string)

	return
}

func (o *FrameCpyToCtx) InitCtx(ref *map[string]interface{}) (err error) {
	o.context = ref

	return
}

func (o *FrameCpyToCtx) Close() (err error) {
	err = o.frame.Close()
	if err != nil {
		return err
	}

	return
}

func (o *FrameCpyToCtx) Process(frame *gocv.Mat) (err error) {
	o.frame.Close()
	o.frame = frame.Clone()

	ctx := *o.context
	ctx[o.Key] = &o.frame

	return
}
