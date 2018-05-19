package server

import (
	"github.com/gorilla/mux"
	"github.com/jurevic/facegrinder/pkg/api"
	"github.com/jurevic/facegrinder/pkg/api/v1/auth"
	"github.com/jurevic/facegrinder/pkg/api/v1/channel"
	"github.com/jurevic/facegrinder/pkg/datastore"
	m "github.com/jurevic/facegrinder/pkg/middleware"
	"github.com/jurevic/facegrinder/pkg/rtmp_server"
	"github.com/rs/cors"
	"log"
	"net/http"
	"github.com/jurevic/facegrinder/pkg/api/v1/processor"
	"github.com/jurevic/facegrinder/pkg/api/v1/face"
)

func Run() {
	auth.Init()
	rtmp_server.Init()
	datastore.Init()
	processor.Init()
	defer datastore.Close()

	r := mux.NewRouter().StrictSlash(true)
	r.Use(m.LoggingMiddleware)

	r.HandleFunc("/api/version", api.GetVersion).Methods("GET")

	// SIGN UP
	r.HandleFunc("/api/v1/users/", auth.Create).Methods("POST")

	// STORAGE
	r.PathPrefix("/storage/").Handler(
		http.StripPrefix("/storage/", http.FileServer(http.Dir("storage"))))

	// AUTH
	ar := r.PathPrefix("/api/auth").Subrouter()
	ar.Methods("POST").Path("/login").HandlerFunc(auth.Login)
	ar.Methods("POST").Path("/refresh").HandlerFunc(auth.Refresh)

	// Sub route for auth protected views
	apr := r.PathPrefix("/api/v1").Subrouter()
	apr.Use(m.AuthMiddleware)

	// USERS
	ur := apr.PathPrefix("/users").Subrouter()
	ur.Methods("GET").Path("/").HandlerFunc(auth.List)
	ur.Methods("GET").Path("/{id}").HandlerFunc(auth.Retrieve)
	ur.Methods("PUT").Path("/{id}").HandlerFunc(auth.Update)
	ur.Methods("DELETE").Path("/{id}").HandlerFunc(auth.Delete)

	// CHANNELS
	cr := apr.PathPrefix("/channels").Subrouter()
	cr.Methods("GET").Path("/").HandlerFunc(channel.List)
	cr.Methods("GET").Path("/{id}").HandlerFunc(channel.Retrieve)
	cr.Methods("GET").Path("/{id}/view").HandlerFunc(channel.View)
	cr.Methods("GET").Path("/{id}/stream").HandlerFunc(channel.Stream)
	cr.Methods("POST").Path("/").HandlerFunc(channel.Create)
	cr.Methods("PUT").Path("/{id}").HandlerFunc(channel.Update)
	cr.Methods("DELETE").Path("/{id}").HandlerFunc(channel.Delete)

	// FACES
	fr := apr.PathPrefix("/faces").Subrouter()
	fr.Methods("GET").Path("/").HandlerFunc(face.List)
	fr.Methods("GET").Path("/{id}").HandlerFunc(face.Retrieve)
	fr.Methods("POST").Path("/").HandlerFunc(face.Create)
	fr.Methods("PUT").Path("/{id}").HandlerFunc(face.Update)
	fr.Methods("DELETE").Path("/{id}").HandlerFunc(face.Delete)

	// PROCESSORS
	pr := apr.PathPrefix("/processors").Subrouter()
	pr.Methods("GET").Path("/").HandlerFunc(processor.List)
	pr.Methods("GET").Path("/choices/").HandlerFunc(processor.ListChoices)
	pr.Methods("GET").Path("/{id}").HandlerFunc(processor.Retrieve)
	pr.Methods("GET").Path("/{id}/run").HandlerFunc(processor.Run)
	pr.Methods("POST").Path("/").HandlerFunc(processor.Create)
	pr.Methods("PUT").Path("/{id}").HandlerFunc(processor.Update)
	pr.Methods("DELETE").Path("/{id}").HandlerFunc(processor.Delete)

	// CORS
	handler := cors.AllowAll().Handler(r)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
