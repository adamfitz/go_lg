package main

import (
	"html/template"
	"log"
	"net/http"
	"net"
	"os/exec"
)


var tmpl = template.Must(template.New("tmpl").ParseFiles("assets/index.html"))


func execCommandHandler(w http.ResponseWriter, r *http.Request) {
	// Print working directory
	output, err := exec.Command("pwd").Output()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(output) // write command output to the response
}


func homePageHandler(w http.ResponseWriter, r *http.Request) {
	// get client ip (removing the port) from the request
	// client ip is displayed at the bottom of index page {{ . }}
	clientIp, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.ExecuteTemplate(w, "index.html", clientIp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func main() {
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/exec", execCommandHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	log.Fatal(http.ListenAndServe("0.0.0.0:9000", nil))
}
