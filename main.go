// Command identity is a fixture service. Generated from testenv/topology/topology.yaml by testenv/topology/generate.js. Do not edit by hand.
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	h := &Handler{svc: NewService(NewStore(), NewClients())}
	r := chi.NewRouter()
	r.Get("/healthz", h.healthz)
	h.register(r)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, r))
}
