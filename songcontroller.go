package songcontroller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"
	"github.com/itchyny/volume-go"
)

// Using a packr box means the html files are bundled up in the binary application.
var templateBox = packr.NewBox("./html")

// tmpl is our pointer to our parsed templates.
var tmpl *template.Template

// This does some initialisation.  It parses our html templates up front.
func init() {

	tmpl = template.New("")
	for _, name := range templateBox.List() {
		t := tmpl.New(name)
		templateStr, err := templateBox.FindString(name)
		if err != nil {
			log.Fatalf("can't find the template in the box: %+v", err)
		}
		template.Must(t.Parse(templateStr))
	}
}

// IndexHandler is the root handler.
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "index.html", nil)

}

// MuteHandler is called to mute the system audio.
func MuteHandler(w http.ResponseWriter, r *http.Request) {
	volume.Mute()
}

// VolumeDownHandler is called to unmute the system audio.
func UnMuteHandler(w http.ResponseWriter, r *http.Request) {
	volume.Unmute()
}

// GetVolumeHandler is called to return the system volume level.
func GetVolumeHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response := make(map[string]string)

	// Get current volume level
	currentVolumeLevel, err := volume.GetVolume()

	if err != nil {
		log.Printf("couldn't retrieve volume level: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		response["volumelevel"] = fmt.Sprint(50)
		response["result"] = "Couldn't parse volume level."
	} else {
		w.WriteHeader(http.StatusOK)
		response["volumelevel"] = fmt.Sprint(currentVolumeLevel)
		response["result"] = "Success"
		json.NewEncoder(w).Encode(response)
	}

}

// SetVolumeHandler is called to set the system volume level to the value passed in.
func SetVolumeHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response := make(map[string]string)

	// Get current volume level
	currentVolumeLevel, err := volume.GetVolume()

	// Extract volumeLevel argument passed in.
	volumeLevelStr := mux.Vars(r)["volumeLevel"]
	volumeLevel, err2 := strconv.Atoi(volumeLevelStr)
	if err != nil || err2 != nil {
		log.Printf("couldn't parse volume level: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		response["volumelevel"] = fmt.Sprint(currentVolumeLevel)
		response["result"] = "Couldn't parse volume level."
		json.NewEncoder(w).Encode(response)
		return
	}

	if volumeLevel < 0 {
		volumeLevel = 0
	}
	if volumeLevel > 100 {
		volumeLevel = 100
	}
	err = volume.SetVolume(volumeLevel)
	if err != nil {
		// SetVolume was not successful
		response["volumelevel"] = fmt.Sprint(currentVolumeLevel)
		response["result"] = "Couldn't set volume level."
	} else {
		w.WriteHeader(http.StatusOK)
		response["volumelevel"] = fmt.Sprint(volumeLevel)
		response["result"] = "Success"
	}
	json.NewEncoder(w).Encode(response)

}
