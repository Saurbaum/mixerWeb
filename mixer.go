package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

type track struct {
	Name string
	ID   int
}

var tracks []track

var mixerVolumeData = "0,0,0,0,0,0,0,0"
var trackVolumeData = "0"
var currentTrack track

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

	for _, audioTrack := range tracks {
		i, err := strconv.Atoi(bodyText)
		if err == nil && i == audioTrack.ID {
			currentTrack = audioTrack
		}
	}

	fmt.Println(currentTrack.Name)
	io.WriteString(w, currentTrack.Name)
}

func pause(w http.ResponseWriter, r *http.Request) {
	fmt.Println("pause")
}

func stop(w http.ResponseWriter, r *http.Request) {
	fmt.Println("stop")
}

func previous(w http.ResponseWriter, r *http.Request) {
	fmt.Println("previous")
	var newID = currentTrack.ID - 1
	if newID < 0 {
		newID = len(tracks) - 1
	}

	for _, audioTrack := range tracks {
		if newID == audioTrack.ID {
			currentTrack = audioTrack
		}
	}

	fmt.Println(currentTrack.Name)
	io.WriteString(w, strconv.Itoa(currentTrack.ID))
}

func next(w http.ResponseWriter, r *http.Request) {
	fmt.Println("next")
	var newID = currentTrack.ID + 1
	if newID >= len(tracks) {
		newID = 0
	}

	for _, track := range tracks {
		if newID == track.ID {
			currentTrack = track
		}
	}

	fmt.Println(currentTrack.Name)
	io.WriteString(w, strconv.Itoa(currentTrack.ID))
}

func trackList(w http.ResponseWriter, r *http.Request) {
	currentTrack = tracks[0]

	fmt.Println(currentTrack.Name)

	b, _ := json.Marshal(tracks)

	fmt.Println(string(b))
	io.WriteString(w, string(b))
}

func main() {
	// Convert to load from file
	tracks = append(tracks, track{"Track 1", 0})
	tracks = append(tracks, track{"Track 2", 1})
	tracks = append(tracks, track{"Track 3", 2})
	tracks = append(tracks, track{"Track 4", 3})
	tracks = append(tracks, track{"Track 5", 4})
	tracks = append(tracks, track{"Track 6", 5})

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("www")))
	mux.HandleFunc("/mixerVolumes", mixerVolumes)
	mux.HandleFunc("/play", play)
	mux.HandleFunc("/pause", pause)
	mux.HandleFunc("/stop", stop)
	mux.HandleFunc("/previous", previous)
	mux.HandleFunc("/next", next)
	mux.HandleFunc("/trackList", trackList)

	// implment this
	//mux.HandleFunc("/filterTracks", filterTracks)
	http.ListenAndServe(":80", mux)
}
