package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Track struct {
	Name string
	Id   int
}

var mixerVolumeData = "0,0,0,0,0,0,0"

func mixerVolumes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println(mixerVolumeData)
		io.WriteString(w, mixerVolumeData)

	case "POST":
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)
		var bodyText = string(body)
		fmt.Println(bodyText)
		io.WriteString(w, bodyText)
		mixerVolumeData = bodyText
	}
}

func play(w http.ResponseWriter, r *http.Request) {
	fmt.Println("play")
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	var bodyText = string(body)
	fmt.Println(bodyText)
	io.WriteString(w, bodyText)
}

func pause(w http.ResponseWriter, r *http.Request) {
	fmt.Println("pause")
}

func stop(w http.ResponseWriter, r *http.Request) {
	fmt.Println("stop")
}

func previous(w http.ResponseWriter, r *http.Request) {
	fmt.Println("previous")
}

func next(w http.ResponseWriter, r *http.Request) {
	fmt.Println("next")
}

func playbackVolume(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	var bodyText = string(body)
	fmt.Println(bodyText)
}

func trackList(w http.ResponseWriter, r *http.Request) {

	var tracks []Track
	tracks = append(tracks, Track{"Track 1", 0})
	tracks = append(tracks, Track{"Track 2", 1})
	tracks = append(tracks, Track{"Track 3", 2})
	tracks = append(tracks, Track{"Track 4", 3})
	tracks = append(tracks, Track{"Track 5", 4})

	b, _ := json.Marshal(tracks)

	fmt.Println(string(b))
	io.WriteString(w, string(b))
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("www")))
	mux.HandleFunc("/mixerVolumes", mixerVolumes)
	mux.HandleFunc("/play", play)
	mux.HandleFunc("/pause", pause)
	mux.HandleFunc("/stop", stop)
	mux.HandleFunc("/previous", previous)
	mux.HandleFunc("/next", next)
	mux.HandleFunc("/playbackVolume", playbackVolume)
	mux.HandleFunc("/trackList", trackList)
	http.ListenAndServe(":8000", mux)
}
