package main

import (
	"log"
	"net/http"

	"github.com/dong-tran/goinaction/uri"
	"github.com/rs/cors"
)

func startServer() {
	log.Printf("Starting server at http://localhost:5508")
	var controllers = uri.CreateControllers()
	log.Printf("Found %d endpoint need to add", len(controllers))
	mux := http.NewServeMux()
	for _, ctr := range controllers {
		mux.HandleFunc(ctr.Path(), ctr.Handle)
		log.Printf("Added path: %s", ctr.Path())
	}
	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":5508", handler))
}

func main() {
	startServer()
}
