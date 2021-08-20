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

	listenAddr := fmt.Sprintf(":%d", 80)
	log.Printf("listening on %s", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, router))

	// vol, err := volume.GetVolume()
	// if err != nil {
	// 	log.Fatalf("get volume failed: %+v", err)
	// }
	// fmt.Printf("current volume: %d\n", vol)

	// err = volume.SetVolume(10)
	// if err != nil {
	// 	log.Fatalf("set volume failed: %+v", err)
	// }
	// fmt.Printf("set volume success\n")

	// err = volume.Mute()
	// if err != nil {
	// 	log.Fatalf("mute failed: %+v", err)
	// }

	// err = volume.Unmute()
	// if err != nil {
	// 	log.Fatalf("unmute failed: %+v", err)
	// }

}
