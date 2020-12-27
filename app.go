package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jakoubek/go-example-webservice/views"
	"log"
	"net/http"
	"os"
)

func main() {
	run()
}

func run() {

	r := mux.NewRouter()
	setupRoutes(r)

	log.Print("Starting server on " + getServerPort())
	http.ListenAndServe(getServerPort(), r)

}

func setupRoutes(r *mux.Router) {
	r.HandleFunc("/", homepage)
	r.HandleFunc("/faq", faq)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	homeView := views.NewView("default", "views/home.gohtml")
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func faq(w http.ResponseWriter, r *http.Request) {
	faqView := views.NewView("default", "views/faq.gohtml")
	w.Header().Set("Content-Type", "text/html")
	must(faqView.Render(w, nil))
}

func rootInfo(w http.ResponseWriter, r *http.Request) {

	type result struct {
		Result string `json:"result"`
		Info   string `json:"info"`
	}

	response := result{
		Result: "OK",
		Info:   "Go to https://www.onetimecode.net for information on how to access the API. See /status for API health.",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

func getServerPort() string {
	if port, ok := os.LookupEnv("PORT"); ok {
		return ":" + port
	}
	return ":3000"
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
