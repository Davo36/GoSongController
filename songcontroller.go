package songcontroller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gobuffalo/packr"
	"github.com/itchyny/volume-go"
)

var volumeIncrementDecrement = 10

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

	type volumeResponse struct {
		Volume       string
		ErrorMessage string
	}

	currentVolume, err := volume.GetVolume()
	errorMessage := ""
	volumeStr := ""
	if err != nil {
		log.Printf("couldn't retrieve volume: %+v", err)
		errorMessage = fmt.Sprint("Couldn't retrieve volume:", err)
	} else {
		volumeStr = strconv.Itoa(currentVolume)
	}

	resp := &volumeResponse{
		Volume:       volumeStr,
		ErrorMessage: errorMessage,
	}

	tmpl.ExecuteTemplate(w, "index.html", resp)
}

// VolumeUpHandler is called to raise the system volume.
func VolumeUpHandler(w http.ResponseWriter, r *http.Request) {
	volume.IncreaseVolume(volumeIncrementDecrement)
}

// VolumeDownHandler is called to lower the system volume.
func VolumeDownHandler(w http.ResponseWriter, r *http.Request) {
	volume.IncreaseVolume(-volumeIncrementDecrement)
}

// MuteHandler is called to mute the system audio.
func MuteHandler(w http.ResponseWriter, r *http.Request) {
	volume.Mute()
}

// VolumeDownHandler is called to unmute the system audio.
func UnMuteHandler(w http.ResponseWriter, r *http.Request) {
	volume.Unmute()
}
