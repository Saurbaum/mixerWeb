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

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("www")))
	mux.HandleFunc("/setVolumes", setVolumes)
	http.ListenAndServe(":8000", mux)
}
