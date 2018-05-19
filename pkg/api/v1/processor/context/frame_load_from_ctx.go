package context

import (
	"errors"
	"gocv.io/x/gocv"
)

type FrameLoadFromCtx struct {
	Key string `json:"key"`

	context *map[string]interface{}
}

func (o *FrameLoadFromCtx) Default() (err error) {
	o.Key = "frame_1"

	return
}

func (o *FrameLoadFromCtx) Init(params map[string]interface{}) (err error) {
	if key, ok := params["key"]; ok {
		o.Key, ok = key.(string)
		if !ok {
			err = errors.New("key is not string type")
			return err
		}
	}

	return
}

func (o *FrameLoadFromCtx) InitCtx(ref *map[string]interface{}) (err error) {
	o.context = ref

	return
}

func (o *FrameLoadFromCtx) Process(frame *gocv.Mat) (err error) {
	ctx := *o.context

	loaded, ok := ctx[o.Key].(*gocv.Mat)
	if !ok {
		err = errors.New("no frame with specified key")
		return err
	}

	loaded.CopyTo(*frame)

	return
}
