package server

import (
	"github.com/gorilla/mux"
	"github.com/jurevic/facegrinder/pkg/api"
	"github.com/jurevic/facegrinder/pkg/api/v1/auth"
	"github.com/jurevic/facegrinder/pkg/api/v1/channel"
	"github.com/jurevic/facegrinder/pkg/api/v1/face"
	"github.com/jurevic/facegrinder/pkg/api/v1/processor"
	"github.com/jurevic/facegrinder/pkg/datastore"
	m "github.com/jurevic/facegrinder/pkg/middleware"
	"github.com/jurevic/facegrinder/pkg/rtmp_server"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func Run() {
	auth.Init()
	rtmp_server.Init()
	processor.Init()
	datastore.Init()
	defer datastore.Close()

	r := mux.NewRouter().StrictSlash(true)
	r.Use(m.LoggingMiddleware)

	r.HandleFunc("/api/version", api.GetVersion).Methods("GET")

	// SIGN UP
	r.HandleFunc("/api/v1/users/", auth.Create).Methods("POST")

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

	// STORAGE
	r.PathPrefix("/storage").Handler(http.StripPrefix("/storage", http.FileServer(http.Dir("storage/"))))

	// FRONTEND
	r.PathPrefix("/static").Handler(http.FileServer(http.Dir("fe/dist/")))
	r.PathPrefix("/").HandlerFunc(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "fe/dist/index.html")
	}))

	// CORS
	handler := cors.AllowAll().Handler(r)

	// Listen and serve
	if viper.GetString("use_tls") == "yes" {
		log.Fatal(
			http.ListenAndServe(
				viper.GetString("http_listen"),
				http.HandlerFunc(redirectTLS)))

		log.Fatal(
			http.ListenAndServeTLS(
				viper.GetString("https_listen"),
				viper.GetString("cert_path"),
				viper.GetString("key_path"),
				handler))
	} else {
		log.Fatal(
			http.ListenAndServe(
				viper.GetString("http_listen"),
				handler))
	}

	log.Fatal(http.ListenAndServe(viper.GetString("http_listen"), handler))
}

func redirectTLS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://" + r.Host + r.RequestURI, http.StatusMovedPermanently)
}
