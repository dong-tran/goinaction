package main

import (
	"log"
	"net/http"

	"github.com/dong-tran/goinaction/uri"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func startServer() {
	log.Printf("Starting server at http://localhost:5508")
	var controllers = uri.CreateControllers()
	log.Printf("Found %d endpoint need to add", len(controllers))
	r := mux.NewRouter()
	for _, ctr := range controllers {
		r.PathPrefix(ctr.Path()).HandlerFunc(ctr.Handle)
		log.Printf("Added path: %s", ctr.Path())
	}
	fshandler := http.FileServer(http.Dir("./public"))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fshandler))
	handler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(":5508", handler))
}

func main() {
	startServer()
}

// func startServerProxy() {
// 	log.Printf("Starting server at http://localhost:5508")
// 	var controllers = uri.CreateControllers()
// 	log.Printf("Found %d endpoint need to add", len(controllers))
// 	r := mux.NewRouter()
// 	for _, ctr := range controllers {
// 		var fnc middleware.HandleGetMethod
// 		r.PathPrefix(ctr.Path()).HandlerFunc(fnc.Handle(ctr.Handle))
// 		log.Printf("Added path: %s", ctr.Path())
// 	}
// 	var proxy middleware.Proxy
// 	fshandler := proxy.Handler(http.FileServer(http.Dir("./assets")))
// 	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fshandler))
// 	handler := cors.Default().Handler(r)
// 	log.Fatal(http.ListenAndServe(":5508", handler))
// }
