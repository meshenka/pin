package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/meshenka/pin"
)

func main() {
	port := flag.String("p", ":8080", "Define the listening port")

	logger := log.New(os.Stdout, "PIN", log.LstdFlags|log.Lshortfile)

	service := pinService{logger}
	flag.Parse()
	r := mux.NewRouter()

	r.HandleFunc("/", PinHandler(service)).Methods("GET")

	logger.Printf("Starting on port %s", *port)

	srv := &http.Server{
		Handler: r,
		Addr:    *port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 6 * time.Second,
		ReadTimeout:  6 * time.Second,
	}

	logger.Fatal(srv.ListenAndServe())
}

//PinHandler serve / by returning a closure
func PinHandler(s PinService) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		pin := s.Generate()

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Cache-Control", "no-cache")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"pin": pin})
	}
}

// PinService service definition
type PinService interface {
	Generate() string
}

type pinService struct {
	logger *log.Logger
}

func (s pinService) Generate() string {
	pin := pin.Generate()
	s.logger.Printf("Service generate pin %s", pin)
	return pin
}
