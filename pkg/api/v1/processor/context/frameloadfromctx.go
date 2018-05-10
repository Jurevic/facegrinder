package context

import (
	"gocv.io/x/gocv"
)

type FrameLoadFromCtx struct {
	Key string

	context *map[string]interface{}
}

func (o *FrameLoadFromCtx) Init(params map[string]interface{}) (err error) {
	o.Key = params["key"].(string)

	return
}

func (o *FrameLoadFromCtx) InitCtx(ref *map[string]interface{}) (err error) {
	o.context = ref

	return
}

func (o *FrameLoadFromCtx) Process(frame *gocv.Mat) (err error) {
	ctx := *o.context

	frame, err = ctx[o.Key].(*gocv.Mat)
	if err != nil {
		return err
	}

	return
}
