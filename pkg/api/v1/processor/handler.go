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
		response.JsonResponse(make([]string, 0), w)
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

	// Requesting user is processor owner
	processor.OwnerId = userId

	// Create processor
	err = db.Insert(&processor)
	if err != nil {
		panic(err)
	}

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

	chain := new(ProcessingChain)

	// Link chain to user
	chain.UserId = userId

	// Initialise chain
	err = chain.Init(processor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer chain.Close()

	chain.Run()

	response.NoContent(w)
}
