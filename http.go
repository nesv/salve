package main

import (
	"net/http"
	"log"
	"github.com/bmizerany/pat"
	"fmt"
)

func StartHttp(laddr string) error {
	// Create a new pattern muxer.
	pmux := pat.New()
	pmux.Post("/nodes/:addr", http.HandlerFunc(httpAddNode))

	// Get the HTTP server listening, on the provided address.
	http.Handle("/", pmux)
	log.Println("http: listening on", laddr)
	return http.ListenAndServe(laddr, nil)
}

/*
Adds a node to the cluster.
*/
func httpAddNode(w http.ResponseWriter, r *http.Request) {
	nodeAddr := r.URL.Query().Get(":addr")
	log.Printf("adding node %q to ring", nodeAddr)
}
