package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Track struct {
	Name string
	Id   int
}

var tracks []Track

var mixerVolumeData = "0,0,0,0,0,0,0"
var currentTrack Track

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

	for _, track := range tracks {
		i, err := strconv.Atoi(bodyText)
		if err == nil && i == track.Id {
			currentTrack = track
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
	var newId = currentTrack.Id - 1
	if newId < 0 {
		newId = len(tracks) - 1
	}

	for _, track := range tracks {
		if newId == track.Id {
			currentTrack = track
		}
	}

	fmt.Println(currentTrack.Name)
	io.WriteString(w, strconv.Itoa(currentTrack.Id))
}

func next(w http.ResponseWriter, r *http.Request) {
	fmt.Println("next")
	var newId = currentTrack.Id + 1
	if newId >= len(tracks) {
		newId = 0
	}

	for _, track := range tracks {
		if newId == track.Id {
			currentTrack = track
		}
	}

	fmt.Println(currentTrack.Name)
	io.WriteString(w, strconv.Itoa(currentTrack.Id))
}

func playbackVolume(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	var bodyText = string(body)
	fmt.Println(bodyText)
}

func trackList(w http.ResponseWriter, r *http.Request) {
	currentTrack = tracks[0]

	fmt.Println(currentTrack.Name)

	b, _ := json.Marshal(tracks)

	fmt.Println(string(b))
	io.WriteString(w, string(b))
}

func main() {
	tracks = append(tracks, Track{"Track 1", 0})
	tracks = append(tracks, Track{"Track 2", 1})
	tracks = append(tracks, Track{"Track 3", 2})
	tracks = append(tracks, Track{"Track 4", 3})
	tracks = append(tracks, Track{"Track 5", 4})

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
