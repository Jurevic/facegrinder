package api

import (
	"github.com/jurevic/facegrinder/pkg/api/v1/helper/response"
	"net/http"
)

const (
	apiVersion = "v0.1a"
)

type Version struct {
	Version string `json:"version"`
}

func GetVersion(w http.ResponseWriter, r *http.Request) {
	response.JsonResponse(Version{Version: apiVersion}, w)
}
