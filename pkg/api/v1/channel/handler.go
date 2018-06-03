package channel

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jurevic/facegrinder/pkg/api/v1/helper/response"
	"github.com/nareix/joy4/av/avutil"
	"github.com/nareix/joy4/format/flv"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
)

type writeFlusher struct {
	httpFlusher http.Flusher
	io.Writer
}

func (wf writeFlusher) Flush() error {
	wf.httpFlusher.Flush()
	return nil
}

func View(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(int)
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	chb.RLock()
	ch := chb.channels[id]
	chb.RUnlock()

	if ch != nil && ch.OwnerId == userId {
		t := template.New("View stream template")

		t, err := t.Parse(viewStreamTemplate)
		if err != nil {
			log.Fatal("Parse: ", err)
			return
		}

		err = t.Execute(w, ch)
		if err != nil {
			log.Fatal("Execute: ", err)
			return
		}
	} else {
		http.NotFound(w, r)
		return
	}
}

func Stream(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(int)
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	chb.RLock()
	ch := chb.channels[id]
	chb.RUnlock()

	if ch != nil && ch.OwnerId == userId {
		w.Header().Set("Content-Type", "video/x-flv")
		w.Header().Set("Transfer-Encoding", "chunked")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(200)

		flusher := w.(http.Flusher)
		flusher.Flush()

		muxer := flv.NewMuxerWriteFlusher(writeFlusher{httpFlusher: flusher, Writer: w})
		cursor := ch.Que.Latest()

		avutil.CopyFile(muxer, cursor)
	} else {
		http.NotFound(w, r)
		return
	}
}

func List(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(int)

	var userChannels []*Channel

	chb.RLock()
	for _, channel := range chb.channels {
		if channel.OwnerId == userId {
			userChannels = append(userChannels, channel)
		}
	}
	chb.RUnlock()

	response.JsonResponse(userChannels, w)
}

func Retrieve(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	chb.RLock()
	ch := chb.channels[id]
	chb.RUnlock()

	if ch == nil {
		http.NotFound(w, r)
		return
	}

	response.JsonResponse(ch, w)
}

func Create(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(int)

	var ch Channel
	err := json.NewDecoder(r.Body).Decode(&ch)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ch.OwnerId = userId

	chb.RLock()
	ch.Id = len(chb.channels) + 1
	chb.channels[ch.Id] = &ch
	chb.RUnlock()

	response.JsonResponse(ch, w)
}

func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateChannel")

	//id := mux.Vars(r)["id"]
}

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteChannel")

	//id := mux.Vars(r)["id"]
}
