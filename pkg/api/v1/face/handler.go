package face

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jurevic/facegrinder/pkg/api/v1/helper/response"
	"github.com/jurevic/facegrinder/pkg/datastore"
	"github.com/satori/go.uuid"
	"image"
	"image/jpeg"
	_ "image/png"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"strconv"
)

func List(w http.ResponseWriter, r *http.Request) {
	db := datastore.DBConn()

	var faces []Face

	err := db.Model(&faces).Select()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	for i := range faces {
		faces[i].Url = "http://localhost:8080/" + faces[i].Url
	}

	response.JsonResponse(faces, w)
}

func Retrieve(w http.ResponseWriter, r *http.Request) {
	db := datastore.DBConn()

	vars := mux.Vars(r)
	pk := vars["id"]

	userId := r.Context().Value("owner_id").(int)

	face := &Face{Id: pk, OwnerId: userId}
	err := db.Select(face)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response.JsonResponse(face, w)
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

	var face Face
	err = json.Unmarshal(data, &face)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Decode image
	dst := "storage/" + strconv.Itoa(userId) + "/faces"
	face.Url, err = uploadImage(face.Url, dst)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Requesting user is face owner
	face.OwnerId = userId

	// Create face
	err = db.Insert(&face)
	if err != nil {
		panic(err)
	}

	response.JsonResponse(face, w)
}

func uploadImage(b64, dst string) (url string, err error) {
	b64data := b64[strings.IndexByte(b64, ',')+1:]

	//Decode to image format
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(b64data))
	img, format, err := image.Decode(reader)
	if err != nil {
		return
	}

	// Ensure path
	if _, err := os.Stat(dst); os.IsNotExist(err) {
		os.MkdirAll(dst, os.ModePerm)
	}

	// Encode from image format to writer
	dst = dst + "/" + uuid.Must(uuid.NewV4()).String() + "." + format
	f, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return
	}

	// Encode
	err = jpeg.Encode(f, img, &jpeg.Options{Quality: 100})
	if err != nil {
		return
	}

	// Set url
	url = dst

	return
}

func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateFace")

	//id := mux.Vars(r)["id"]
}

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteFace")

	//id := mux.Vars(r)["id"]
}
