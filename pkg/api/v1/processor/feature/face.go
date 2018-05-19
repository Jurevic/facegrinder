package feature

import (
	"fmt"
	"github.com/Kagami/go-face"
	apiFace "github.com/jurevic/facegrinder/pkg/api/v1/face"
	"github.com/jurevic/facegrinder/pkg/api/v1/processor/draw"
	"github.com/jurevic/facegrinder/pkg/datastore"
	"gocv.io/x/gocv"
	"errors"
)

type RecogniseFaces struct {
	Key       string  `json:"key"`
	Jitter    int     `json:"jitter"`
	Threshold float64 `json:"threshold"`
	Skip      int     `json:"skip"`

	faceRec *face.Recognizer
	faces   []face.Face
	matches []int
	boxes   []draw.LabeledBox
	counter int
	context *map[string]interface{}
}

func (o *RecogniseFaces) Default() (err error) {
	o.Key = "faces"
	o.Jitter = 0
	o.Threshold = 0.6
	o.Skip = 10

	return
}

func (o *RecogniseFaces) Init(params map[string]interface{}) (err error) {
	if key, ok := params["key"]; ok {
		o.Key, ok = key.(string)
		if !ok {
			err = errors.New("key is not string type")
			return err
		}
	}
	if jitter, ok := params["jitter"]; ok {
		o.Jitter, ok = jitter.(int)
		if !ok {
			err = errors.New("jitter is not int type")
			return err
		}
	}
	if threshold, ok := params["threshold"]; ok {
		o.Threshold, ok = threshold.(float64)
		if !ok {
			err = errors.New("threshold is not float type")
			return err
		}
	}
	if skip, ok := params["skip"]; ok {
		o.Skip, ok = skip.(int)
		if !ok {
			err = errors.New("skip is not int type")
			return err
		}
	}

	o.faceRec, err = face.NewRecognizer("models", o.Threshold, uint(o.Jitter))
	if err != nil {
		return err
	}

	//o.initKnownFaces(userId)

	return
}

func (o *RecogniseFaces) InitCtx(ref *map[string]interface{}) (err error) {
	o.context = ref

	return
}

func (o *RecogniseFaces) Process(frame *gocv.Mat) (err error) {
	if o.counter == 0 {
		o.getFaces(frame)
		if err != nil {
			return err
		}

		o.classifyFaces()
		if err != nil {
			return err
		}

		// Create labeled boxes
		o.boxes = make([]draw.LabeledBox, len(o.faces))
		for i := range o.faces {
			o.boxes[i] = draw.LabeledBox{
				Rectangle: o.faces[i].Rectangle,
				Label:     "test",
			}
		}

		// Output boxes to context
		ctx := *o.context
		ctx[o.Key] = &o.boxes
	}

	if o.counter >= o.Skip {
		o.counter = 0
	} else {
		o.counter++
	}

	return
}

func (o *RecogniseFaces) initKnownFaces(userId int) {
	db := datastore.DBConn()

	var faces []apiFace.Face
	err := db.Model(&faces).Where("owner_id = ?", userId).Select()
	if err != nil {
		return
	}

	var descriptors []face.Descriptor
	for i := range faces {
		knownFace, err := o.faceRec.RecognizeSingleFile(faces[i].Url)
		if err != nil {
			fmt.Println(err)
			return
		}
		descriptors = append(descriptors, knownFace.Descriptor)
	}

	o.faceRec.SetSamples(descriptors)
}

func (o *RecogniseFaces) getFaces(mat *gocv.Mat) (err error) {
	o.faces, err = o.faceRec.RecognizeMat(mat.ToBytes(), mat.Rows(), mat.Cols())
	if err != nil {
		return err
	}

	return
}

func (o *RecogniseFaces) classifyFaces() (err error) {
	var matches []int
	for i := range o.faces {
		matches = append(
			matches,
			o.faceRec.Classify(o.faces[i].Descriptor))
	}

	o.matches = matches

	return
}
