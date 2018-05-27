package processor

import (
	"github.com/jurevic/facegrinder/pkg/datastore"
	"encoding/json"
)

func CreateDefault(userId int) {
	db := datastore.DBConn()

	for _, chain := range defaultChains {
		processor := new(Processor)

		// Unmarshal default processor
		err := json.Unmarshal([]byte(chain), processor)
		if err != nil {
			panic(err)
		}

		// Requesting user is processor owner
		processor.OwnerId = userId

		// Create processor
		err = db.Insert(processor)
		if err != nil {
			panic(err)
		}
	}
}

var defaultChains = []string{
	defaultCamChain,
	defaultRTMPChain,
	defaultRTMPNoOpChain,
}

const defaultCamChain = `{
	"name": "Default",
	"nodes": [
		{
			"key": "input_cam",
			"params": {}
		},
		{
			"key": "cpy_to_ctx",
			"params": {}
		},
		{
			"key": "resize",
			"params": {
				"scale": 0.5
			}
		},
		{
			"key": "rgba_to_bgr",
			"params": {}
		},
		{
			"key": "face_recogniser",
			"params": {}
		},
		{
			"key": "load_from_ctx",
			"params": {}
		},
		{
			"key": "label_faces",
			"params": {
				"scale": 2
			}
		},
		{
			"key": "output_imshow",
			"params": {}
		}
	]
}`

const defaultRTMPChain = `{
	"name": "Default RTMP",
	"nodes": [
		{
			"key": "input_rtmp",
			"params": {}
		},
		{
			"key": "cpy_to_ctx",
			"params": {}
		},
		{
			"key": "resize",
			"params": {
				"scale": 0.5
			}
		},
		{
			"key": "rgba_to_bgr",
			"params": {}
		},
		{
			"key": "face_recogniser",
			"params": {}
		},
		{
			"key": "load_from_ctx",
			"params": {}
		},
		{
			"key": "label_faces",
			"params": {
				"scale": 2
			}
		},
		{
			"key": "output_imshow",
			"params": {}
		}
	]
}`

const defaultRTMPNoOpChain = `{
	"name": "Default RTMP NoOp",
	"nodes": [
		{
			"key": "input_rtmp",
			"params": {}
		},
		{
			"key": "output_imshow",
			"params": {}
		}
	]
}`
