package processor

import (
	"encoding/json"
	"fmt"
	"github.com/jurevic/facegrinder/pkg/api/v1/helper/response"
	"github.com/jurevic/facegrinder/pkg/datastore"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
)

func List(w http.ResponseWriter, r *http.Request) {
	db := datastore.DBConn()

	userId := r.Context().Value("user_id").(int)

	var processors []Processor

	err := db.Model(&processors).Where("owner_id = ?", userId).Select()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response.JsonResponse(processors, w)
}

func Retrieve(w http.ResponseWriter, r *http.Request) {
	db := datastore.DBConn()

	vars := mux.Vars(r)
	pk := vars["id"]

	userId := r.Context().Value("user_id").(int)

	processor := &Processor{Id: pk, OwnerId: userId}
	err := db.Select(processor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response.JsonResponse(processor, w)
}

func Create(w http.ResponseWriter, r *http.Request) {
	db := datastore.DBConn()

	userId := r.Context().Value("user_id").(int)

	// Read data
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var processor Processor
	err = json.Unmarshal(data, &processor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Requesting user is face owner
	processor.OwnerId = userId

	// Create face
	err = db.Insert(&processor)
	if err != nil {
		panic(err)
	}

	//var chain ProcessingChain
	//
	//chain.Processors = make([]FrameProcessor, len(processor))
	//
	//for i := range processors {
	//	defaultProcessor := ProcessorsMap[processors[i].Type]
	//	newProcessor := deepcopy.Copy(defaultProcessor)
	//
	//	ini, ok := newProcessor.(Initializer)
	//	if ok {
	//		err = ini.Init(processors[i].Params)
	//		if err != nil {
	//			return
	//		}
	//	}
	//}

	response.JsonResponse(processor, w)
}

func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateProcessingChain")

	//id := mux.Vars(r)["id"]
}

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteProcessingChain")

	//id := mux.Vars(r)["id"]
}

func ListChoices(w http.ResponseWriter, r *http.Request) {
	response.JsonResponse(ProcessorsMap, w)
}

func Run(w http.ResponseWriter, r *http.Request) {
	//userId := r.Context().Value("user_id").(int)

	//var chain ProcessingChain

	//chain.Input = new(input.Camera)

	//chain.Processors = make([]FrameProcessor, 2)
	//chain.Processors[0] = new(stats.Fps)
	//chain.Processors[1] = new(output.IMShow)
	//
	//params := make(map[string]interface{})
	//
	//params["scale"] = 2.0
	//params["name"] = "Test"
	//
	//params["user_id"] = 1
	//params["jitter"] = 1
	//params["threshold"] = 0.6
	//params["refresh_interval"] = 1
	//params["font_size"] = 2
	//params["font_thickness"] = 1
	//params["box_thickness"] = 1
	//params["x"] = 100
	//params["y"] = 100
	//params["color"] = "R:255 G:0 B:0 A:0"
	//
	//params["key"] = "frame"
	//
	//chain.Init(params)
	//chain.Run()
	//defer chain.Close()

	//InitFacePredictor()
	//InitKnownFaces(userId)
	//ProcessFromRtmpChannel()
	// ProcessFromCam()

	response.NoContent(w)
}
