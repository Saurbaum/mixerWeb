package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func setVolumes(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	var bodyText = string(body)
	fmt.Println(bodyText)
	io.WriteString(w, bodyText)
}

func play(w http.ResponseWriter, r *http.Request) {
	fmt.Println("play")
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

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("www")))
	mux.HandleFunc("/setVolumes", setVolumes)
	mux.HandleFunc("/play", play)
	mux.HandleFunc("/pause", pause)
	mux.HandleFunc("/stop", stop)
	mux.HandleFunc("/previous", previous)
	mux.HandleFunc("/next", next)
	http.ListenAndServe(":8000", mux)
}
