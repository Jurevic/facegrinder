package handler

import (
	"net/http"
)

type ModelCrudHandler interface {
	getModel() interface{}
	getQuery()
}

func ListHandler(w http.ResponseWriter, _ *http.Request) {
	//users := FindAll()
	//
	//bytes, err := json.Marshal(users)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//}t
	//
	//writeJsonResponse(w, bytes)
}

func RetrieveHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//id := vars["id"]
	//
	//user := new(User)
	//
	//user, ok := user.Find(id)
	//if !ok {
	//	http.NotFound(w, r)
	//	return
	//}
	//
	//response.JsonResponse(user, w)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	//body, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//}
	//
	//user := new(User)
	//err = json.Unmarshal(body, user)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//}
	//
	//User.Save(user.Name, user)
	//
	//w.Header().Set("Location", r.URL.Path+"/"+user.Name)
	//w.WriteHeader(http.StatusCreated)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//name := vars["name"]
	//
	//body, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//}
	//
	//user := new(User)
	//err = json.Unmarshal(body, user)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//}
	//
	//User.Save(name, user)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//id := vars["id"]
	//
	//user := new(User)
	//
	//ok := user.Remove(id)
	//if !ok {
	//	http.NotFound(w, r)
	//	return
	//}
	//
	//w.WriteHeader(http.StatusNoContent)
}
