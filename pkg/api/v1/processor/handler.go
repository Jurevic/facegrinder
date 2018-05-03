package processor

import (
	"net/http"
	"fmt"
)

func List(w http.ResponseWriter, r *http.Request) {
	//userId := r.Context().Value("user_id").(string)
	//
	//var userChannels []*Channel

	InitFacePredictor()
	InitKnownFaces()
	// processor.ProcessFromRtmpChannel()
	ProcessFromCam()

	//response.JsonResponse(userChannels, w)
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
	fmt.Println("UpdateChannel")

	//id := mux.Vars(r)["id"]
}

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteChannel")

	//id := mux.Vars(r)["id"]
}