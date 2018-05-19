package context

import (
	"errors"
	"gocv.io/x/gocv"
)

type FrameCpyToCtx struct {
	Key string `json:"key"`

	frame   gocv.Mat
	context *map[string]interface{}
}

func (o *FrameCpyToCtx) Default() (err error) {
	o.Key = "frame_1"

	return
}

func (o *FrameCpyToCtx) Init(params map[string]interface{}) (err error) {
	if key, ok := params["key"]; ok {
		o.Key, ok = key.(string)
		if !ok {
			err = errors.New("key is not string type")
			return err
		}
	}

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
