package main

import (
	"html/template"
	"log"
	"net"
	"net/http"
)

var tmpl = template.Must(template.New("tmpl").ParseFiles("assets/index.html"))

func execActionCommandHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the selected query type
	query := r.FormValue("query")

	// Route to the corresponding function based on query value
	switch query {
	case "bgp":
		handleBGP(w, r)
	case "summary":
		handleBGPSummary(w, r)
	case "ping":
		handlePing(w, r)
	case "trace":
		handleTrace(w, r)
	default:
		http.Error(w, "Invalid query type", http.StatusBadRequest)
	}
}

func handleBGP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("BGP handler executed"))
}

func handleBGPSummary(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("BGP Summary handler executed"))
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ping handler executed"))
}

func handleTrace(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Trace handler executed"))
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
	http.HandleFunc("/execAction", execActionCommandHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	log.Fatal(http.ListenAndServe("0.0.0.0:9000", nil))
}
