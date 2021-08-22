package main

import (
	songcontroller "GoSongController"
	"fmt"
	"log"
	"net/http"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	// Serve up static content.
	static := packr.NewBox("../../static")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(static)))

	// UI handlers.
	router.HandleFunc("/", songcontroller.IndexHandler).Methods("GET")
	router.HandleFunc("/mute", songcontroller.MuteHandler).Methods("GET")
	router.HandleFunc("/unmute", songcontroller.UnMuteHandler).Methods("GET")
	router.HandleFunc("/getvolume", songcontroller.GetVolumeHandler).Methods("GET")
	router.HandleFunc("/setvolume/{volumeLevel:[0-9]+}", songcontroller.SetVolumeHandler).Methods("POST")

	listenAddr := fmt.Sprintf(":%d", 80)
	log.Printf("listening on %s", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, router))
}
