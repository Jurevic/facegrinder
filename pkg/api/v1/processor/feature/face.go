package feature

import (
	"github.com/jurevic/facegrinder/pkg/datastore"
	"github.com/Kagami/go-face"
	"fmt"
	"gocv.io/x/gocv"
	apiFace "github.com/jurevic/facegrinder/pkg/api/v1/face"
	"github.com/jurevic/facegrinder/pkg/api/v1/processor/draw"
)

type FaceRecogniser struct {
	Key string

	faceRec *face.Recognizer
	faces []face.Face
	matches []int
	boxes []draw.LabeledBox
	context *map[string]interface{}
}

func (o *FaceRecogniser) Init(params map[string]interface{}) (err error) {
	userId := params["user_id"].(int)
	o.Key = params["key"].(string)
	jitter := params["jitter"].(int)
	threshold := params["threshold"].(float64)

	o.faceRec, err = face.NewRecognizer("models", threshold, uint(jitter))
	if err != nil {
		return err
	}

	o.initKnownFaces(userId)

	return
}

func (o *FaceRecogniser) InitCtx(ref *map[string]interface{}) (err error) {
	o.context = ref

	return
}

func (o *FaceRecogniser) Process(frame *gocv.Mat) (err error) {
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
			Label: "test",
		}
	}

	// Output boxes to context
	ctx := *o.context
	ctx[o.Key] = &o.boxes

	return
}

func (o *FaceRecogniser) initKnownFaces(userId int) {
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

func (o *FaceRecogniser) getFaces(mat *gocv.Mat) (err error) {
	o.faces, err = o.faceRec.RecognizeMat(mat.ToBytes(), mat.Rows(), mat.Cols())
	if err != nil {
		return err
	}

	return
}

func (o *FaceRecogniser) classifyFaces() (err error) {
	var matches []int
	for i := range o.faces {
		matches = append(
			matches,
			o.faceRec.Classify(o.faces[i].Descriptor))
	}

	o.matches = matches

	return
}
