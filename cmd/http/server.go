package main

import (
	"os"
    "context"
	"github.com/meshenka/pin/cmd"
	"github.com/meshenka/pin/transport/http"
)

func main() {
	if err := cmd.Run(run); err != nil {
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	return http.Listen(
		ctx,
		http.WithEndpoint(
			cmd.Env("APPLICATION_ADDR", ":8080"),
			cmd.Env("HTTP_TIMEOUT", "5s"),
		),
		http.WithLogLevel(
			cmd.Env("LOGLEVEL", "debug"),
		),
	)
}

//func main() {
//	port := flag.String("p", ":8080", "Define the listening port")

//	logger := log.New(os.Stdout, "PIN", log.LstdFlags|log.Lshortfile)

//	service := pin.NewGenerator()
//	flag.Parse()
//	r := mux.NewRouter()

//	r.HandleFunc("/", PinHandler(service)).Methods("GET")

//	logger.Printf("Starting on port %s", *port)

//	srv := &http.Server{
//		Handler: r,
//		Addr:    *port,
//		// Good practice: enforce timeouts for servers you create!
//		WriteTimeout: 6 * time.Second,
//		ReadTimeout:  6 * time.Second,
//	}

//	logger.Fatal(srv.ListenAndServe())
//}

//type Generator interface {
//	Generate() string
//}

////PinHandler serve / by returning a closure
//func PinHandler(g Generator) func(w http.ResponseWriter, r *http.Request) {

//	return func(w http.ResponseWriter, r *http.Request) {
//		pin := g.Generate()

//		w.Header().Set("Content-Type", "application/json; charset=utf-8")
//		w.Header().Set("Cache-Control", "no-cache")

//		w.WriteHeader(http.StatusOK)
//		json.NewEncoder(w).Encode(map[string]string{"pin": pin})
//	}
//}
