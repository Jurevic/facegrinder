package processor

import (
	"net/http"
	"fmt"
	"github.com/jurevic/facegrinder/pkg/api/v1/processor/input"
	"github.com/jurevic/facegrinder/pkg/api/v1/helper/response"
	"github.com/jurevic/facegrinder/pkg/api/v1/processor/output"
	"github.com/jurevic/facegrinder/pkg/api/v1/processor/stats"
	"github.com/jurevic/facegrinder/pkg/api/v1/processor/feature"
	"github.com/jurevic/facegrinder/pkg/api/v1/processor/draw"
)

func List(w http.ResponseWriter, r *http.Request) {
	//userId := r.Context().Value("user_id").(int)


	var chain ProcessingChain

	chain.Input = new(input.Camera)

	chain.Processors = make([]FrameProcessor, 4)
	chain.Processors[0] = new(stats.Fps)
	chain.Processors[1] = new(feature.FaceRecogniser)
	chain.Processors[2] = new(draw.LabeledBoxes)
	chain.Processors[3] = new(output.IMShow)

	params := make(map[string]interface{})

	params["scale"] = 2.0
	params["name"] = "Test"

	params["user_id"] = 1
	params["jitter"] = 1
	params["threshold"] = 0.6
	params["refresh_interval"] = 1
	params["font_size"] = 2
	params["font_thickness"] = 1
	params["box_thickness"] = 1
	params["x"] = 100
	params["y"] = 100
	params["color"] = "R:255 G:0 B:0 A:0"

	params["key"] = "frame"

	chain.Init(params)
	chain.Run()
	defer chain.Close()

	//InitFacePredictor()
	//InitKnownFaces(userId)
	//ProcessFromRtmpChannel()
	// ProcessFromCam()

	response.NoContent(w)
}



func Retrieve(w http.ResponseWriter, r *http.Request) {
	//id := mux.Vars(r)["id"]
	//userId := r.Context().Value("user_id").(string)
	//
	//response.JsonResponse(ch, w)
}

func Create(w http.ResponseWriter, r *http.Request) {
	//userId := r.Context().Value("user_id").(string)
	//
	//response.JsonResponse(ch, w)
}

func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateProcessor")

	//id := mux.Vars(r)["id"]
}

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteChannel")

	//id := mux.Vars(r)["id"]
}